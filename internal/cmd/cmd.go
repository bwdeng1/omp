package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"omp/internal/controller/common"
	"omp/internal/controller/permission"
	"omp/internal/controller/public"
	"omp/internal/controller/role"
	"omp/internal/controller/user"
	"omp/internal/service"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					public.NewV1(),
				)
				//权限控制路由
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(service.Middleware().Auth)
					group.Bind(
						//系统管理
						user.NewV1(),       // 用户
						permission.NewV1(), // 权限
						role.NewV1(),       // 角色
						common.NewV1(),     // 公共、通用接口
					)

				})

			})
			s.Run()
			return nil
		},
	}
)
