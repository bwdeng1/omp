package user

import (
	"context"
	"github.com/gogf/gf/v2/util/gutil"
	"omp/api"
	"omp/internal/dao"
	"omp/internal/model/comb"
	"omp/internal/model/do"
	"omp/internal/model/entity"
	"omp/internal/service"
	"omp/utility/util"
)

type sUser struct{}

var cols = dao.User.Columns()

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

func (s *sUser) Add(ctx context.Context, in *entity.User) (err error) {
	// 1. 密码加密
	in.Password, err = util.EncryptPassword(in.Password)
	if err != nil {
		return
	}
	// 2. 调用 DAO 将数据插入数据库
	_, err = dao.User.Ctx(ctx).Insert(in)
	return
}

func (*sUser) Upt(ctx context.Context, in *do.User) (err error) {
	if !gutil.IsEmpty(in.Password) {
		in.Password, err = util.EncryptPassword(in.Password.(string))
		if err != nil {
			return
		}
	}
	_, err = dao.User.Ctx(ctx).WherePri(in.Id).OmitNilData().Data(in).Update()
	return
}

func (*sUser) GetLst(ctx context.Context) (out []*entity.User, err error) {
	err = dao.User.Ctx(ctx).FieldsEx(cols.RoleIds, cols.Phone, cols.Password, cols.DeptId).OrderDesc(cols.Id).Scan(&out)
	return
}

func (*sUser) GetPageLst(ctx context.Context, in *api.PageLstReq) (out *api.PageLstRes[*entity.User], err error) {
	out = &api.PageLstRes[*entity.User]{}
	// 1. 构建基础查询模型
	m := dao.User.Ctx(ctx).Safe(true)
	// 2. 处理动态搜索条件
	//如果 Search 不为空，就对 username 和 real_name 做模糊匹配（LIKE），任意一个匹配就算符合。
	//.WhereOr 表示追加一个 OR 条件 到 SQL 的 WHERE 子句里
	//m.Builder() 是生成一个新的 Where 构造器（builder），用来组合复杂的条件
	if !gutil.IsEmpty(in.Search) {
		m = m.WhereOr(m.Builder().WhereOrLike(cols.Username, in.SearchStr()).WhereOrLike(cols.RealName, in.SearchStr()))
	}

	//in.Wheres 是一个 *gjson.Json，代表前端传来的 JSON 格式的【筛选条件】。
	//.Get("xxx") 是从 JSON 里找对应字段。
	//.IsNil() 判断有没有传。
	//.Bool() / .Int() 是把 JSON 里的值转成 Go 的布尔值或整数。
	if enabled := in.Wheres.Get("enabled"); !enabled.IsNil() {
		m = m.Where(cols.Enabled, enabled.Bool())
	}

	if deptId := in.Wheres.Get("deptId"); !deptId.IsNil() {
		m = m.Where(cols.DeptId, deptId.Int())
	}
	// 3. 执行查询，同时完成分页和总数统计
	err = m.Offset(in.Offset()).Limit(in.Limit()).FieldsEx(cols.Password).
		ScanAndCount(&out.List, &out.Total, false)
	return
}

func (*sUser) Get(ctx context.Context, userDo *do.User) (out *entity.User, err error) {
	err = dao.User.Ctx(ctx).Where(userDo).OmitNilWhere().Limit(1).Scan(&out)
	return
}
func (*sUser) GetComb(ctx context.Context, userDo *do.User) (out *comb.User, err error) {
	if err = dao.User.Ctx(ctx).Where(userDo).OmitNilWhere().Limit(1).Scan(&out); err != nil {
		return
	}
	if err = dao.Role.Ctx(ctx).WhereIn(dao.Role.Columns().Id, out.RoleIds.Array()).Scan(&out.Roles); err != nil {
		return
	}
	return
}

func (*sUser) GetCombLst(ctx context.Context) (out []*comb.User, err error) {
	if err = dao.User.Ctx(ctx).Where(cols.Enabled, true).Scan(&out); err != nil {
		return
	}

	for _, user := range out {
		if err = dao.Role.Ctx(ctx).WhereIn(dao.Role.Columns().Id, user.RoleIds.Array()).Scan(&user.Roles); err != nil {
			return
		}
	}
	return
}

func (*sUser) Del(ctx context.Context, in *do.User) (err error) {
	_, err = dao.User.Ctx(ctx).Where(in).OmitNilWhere().Delete()
	return
}
