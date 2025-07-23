package context

import (
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/gogf/gf/v2/frame/g"
	"omp/internal/consts"
	"omp/internal/model/comb"
	"omp/internal/service"
	"strings"
)

// 填写规则书
func (*sContext) initCasbin(ctx context.Context) (*casbin.Enforcer, error) {
	g.Log().Debug(ctx, "开始初始化 casbin 执行器。")
	//创建了一个 Casbin 的决策者（Enforcer），它已经知道了规则（模型），但里面一条具体的策略都没有
	enforcer, err := casbin.NewEnforcer(newCasbinModel())
	if err != nil {
		return nil, err
	}
	//从数据库获取所有角色及其关联的权限树
	roleList, err := service.Role().GetCombList(ctx)
	if err != nil {
		return nil, err
	}
	//给这个函数类型起了个名字
	type recursiveHandlePermissionsFunc func(roleCode string, permissions []*comb.Permission) error
	//声明一个变量，变量的类型就是 recursiveHandlePermissionsFunc
	var recursiveHandlePermissions recursiveHandlePermissionsFunc
	//递归地处理角色的权限树，然后把每条权限加到 Casbin 里
	recursiveHandlePermissions = func(roleCode string, permissions []*comb.Permission) error {
		for _, permission := range permissions {
			//判断是否是可执行目录
			if permission.Type == consts.PERMISSION_TYPE_ABLE {
				//.Array() 就是把原本存放在 JSON 里的数组，转换成 Go 能直接用的切片
				for _, backendRouteInfo := range permission.BRoutes.Array() {
					//分成两段
					routeInfoArr := strings.SplitN(backendRouteInfo.(string), ":", 2)
					//strings.ToUpper 是把字符串转换成大写字母的函数
					method := strings.ToUpper(routeInfoArr[0])
					routePath := routeInfoArr[1]
					//判断是否已经存在策略，存在就跳过
					if enforcer.HasPolicy(roleCode, routePath, method) {
						continue
					}
					if _, err := enforcer.AddPolicy(roleCode, routePath, method); err != nil {
						return err
					}
					g.Log().Debugf(ctx, "成功添加策略: %s, %s, %s", roleCode, routePath, method)
				}
			} else {
				//permission.Children 是当前权限的下级权限列表  如果有，就要继续递归处理
				if len(permission.Children) > 0 {
					if err = recursiveHandlePermissions(roleCode, permission.Children); err != nil {
						return err
					}
				}
			}
		}
		return nil
	}
	//每个角色都调用一下
	for _, role := range roleList {
		if err := recursiveHandlePermissions(role.Code, role.Permissions); err != nil {
			return nil, err
		}
	}
	//用户服务中获取用户列表
	userList, err := service.User().GetCombLst(ctx)
	if err != nil {
		return nil, err
	}

	for _, user := range userList {
		//存放当前用户所有角色的编码
		var roleCodes []string
		//用户的角色列表，将每个角色的 Code（角色标识符）追加到 roleCodes
		for _, role := range user.Roles {
			roleCodes = append(roleCodes, role.Code)
		}
		//调用 Casbin 的 AddRolesForUser，为用户绑定这些角色
		if _, err := enforcer.AddRolesForUser(user.Username, roleCodes); err != nil {
			return nil, err
		}
		g.Log().Debugf(ctx, "成功添加角色: %s, %v", user.Username, roleCodes)
	}
	g.Log().Debug(ctx, "初始化 casbin 执行器成功！")
	return enforcer, nil
}

// 定义规则书
func newCasbinModel() model.Model {
	m := model.NewModel()
	//请求定义
	m.AddDef("r", "r", "sub, obj, act")
	//策略定义
	m.AddDef("p", "p", "sub, obj, act")
	//分组/角色继承
	m.AddDef("g", "g", "_, _")
	//效果
	m.AddDef("e", "e", "some(where (p.eft == allow))")
	m.AddDef("m", "m", `g(r.sub, p.sub) && keyMatch2(r.obj, p.obj) && regexMatch(r.act, p.act) || g(r.sub, "admin")`)
	return m
}
