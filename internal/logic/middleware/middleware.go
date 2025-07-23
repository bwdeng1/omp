package middleware

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"net/http"
	"omp/internal/service"
)

type sMiddleware struct{}

// init() 在程序启动时自动执行，把这个中间件实例注册到全局服务里 ➜
// 其他地方可以随时通过 service.Middleware() 拿到它来用。
func init() {
	// 1. 自动注册
	// 在程序启动时，通过 service.RegisterMiddleware 将当前间件的实例注册到全局服务中
	// 这样其他地方就可以通过 service.Middleware() 来获取和使用它。
	service.RegisterMiddleware(New())
}

// init() 调用 New() ➜ 创建一个 sMiddleware 的实例
// New() 是创建中间件实例的工厂函数
func New() *sMiddleware {
	return &sMiddleware{}
}

func (*sMiddleware) Auth(r *ghttp.Request) {
	// 2. JWT 认证 (Authentication)
	//service.Auth() 获取一个Auth服务实例 .MiddlewareFunc() 返回一个 gf框架的中间件函数
	//()(r) 直接执行这个中间件，对当前请求 r 做 JWT 校验
	service.Auth().MiddlewareFunc()(r)
	var (
		requestPath                   = r.Request.URL.Path
		method                        = r.Request.Method
		refreshPermissionPathPrefixes = []string{"/permission", "/role", "/user"}
	)
	//Casbin 权限校验  Enforce() 检查：谁（Username）对哪个接口（requestPath）用什么方法（GET / POST / PUT / DELETE）是否有权限
	//service.CurrentUser(ctx) 就是一个封装好方法：从 Context 里提取 user 信息（比如 id, username, roles, etc）。
	//service.Context() ➜ 拿到封装好的上下文服务 .Ctx() ➜ 里面挂着全局的 CasbinEnforcer 实例
	//CasbinEnforcer ➜ 是 Casbin 提供的核心鉴权引擎，封装了：
	//加载了模型（model.conf）
	//加载了策略（policy）
	//提供了 Enforce 方法去做匹配判断
	//Enforce(subject, obj, act)
	pass, err := service.Context().Ctx().CasbinEnforcer.Enforce(service.CurrentUser(r.GetCtx()).Username, requestPath, method)
	if err != nil {
		r.Response.WriteJson(g.Map{
			"code":    http.StatusForbidden,
			"message": err,
		})
		r.ExitAll()
	}

	if !pass {
		r.Response.WriteJson(g.Map{
			"code":    http.StatusForbidden,
			"message": "没有接口权限",
		})
		r.ExitAll()
	}

	r.Middleware.Next()
	//对非 GET 请求，如果请求路径以某些前缀开头（比如涉及到权限、角色、用户的接口）
	//在接口执行后就自动刷新 Casbin 的权限策略缓存，以保证最新的权限生效。
	if method != "GET" {
		for _, prefix := range refreshPermissionPathPrefixes {
			//gstr.HasPrefix：判断请求路径是否以 prefix 开头
			if gstr.HasPrefix(requestPath, prefix) {
				if err := service.Context().RefreshCasbin(r.GetCtx()); err != nil {
					r.Response.WriteJson(g.Map{
						"code":    http.StatusInternalServerError,
						"message": gerror.Newf("刷新权限失败: %v", err),
					})
					r.ExitAll()
				}
				break
			}
		}
	}

}
