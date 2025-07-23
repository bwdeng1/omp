package service

import (
	"context"
	v1 "omp/api/public/v1"
	"omp/internal/model"
	"omp/internal/model/do"
	"omp/internal/model/entity"
	"omp/utility/util"
	"time"

	"github.com/gogf/gf/v2/encoding/gjson"

	jwt "github.com/gogf/gf-jwt/v2"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

var authService *jwt.GfJWTMiddleware

func Auth() *jwt.GfJWTMiddleware {
	return authService
}

// 信息提取和格式化工具
// gjson.New会把它包装成一个 gjson 对象
// .Scan(&u)这是最神奇的一步。它相当于把从保安那里拿到的零散信息（map）
// 自动地填写到一张标准的名片（model.RequestUser 结构体）里。
func CurrentUser(ctx context.Context) (u *model.RequestUser) {
	_ = gjson.New(authService.GetPayload(ctx)).Scan(&u)
	return
}

func init() {
	ctx := context.Background()
	auth := jwt.New(&jwt.GfJWTMiddleware{
		Realm:           "omp",
		Key:             g.Cfg().MustGet(context.Background(), "jwt.secret").Bytes(),
		Timeout:         g.Cfg().MustGet(ctx, "jwt.expire").Duration(),
		MaxRefresh:      g.Cfg().MustGet(ctx, "jwt.expire").Duration(),
		IdentityKey:     "userId",
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
		Authenticator:   Authenticator,
		Unauthorized:    Unauthorized,
		PayloadFunc:     PayloadFunc,
		IdentityHandler: IdentityHandler,
	})
	authService = auth
}

func PayloadFunc(data any) jwt.MapClaims {
	// 拿到认证官传过来的用户对象
	userInfo := data.(*entity.User)
	// 创建一个空的信息表
	claims := make(jwt.MapClaims)
	// 把需要的信息登记到表上
	claims["userId"] = userInfo.Id
	claims["username"] = userInfo.Username
	claims["realName"] = userInfo.RealName
	return claims
}

// 一个“指针”，专门用来告诉 JWT 框架，在所有用户信息中，哪一个才是代表用户唯一身份的“主键”。
func IdentityHandler(ctx context.Context) interface{} {
	claims := jwt.ExtractClaims(ctx)
	return claims[authService.IdentityKey]
}

func Unauthorized(ctx context.Context, code int, message string) {
	r := g.RequestFromCtx(ctx)
	g.Log().Debug(ctx, message)
	r.Response.WriteJson(g.Map{
		"code":    code,
		"message": message,
	})
	//ExitAll 是 GF 的语法糖
	r.ExitAll()
}

// Authenticator 用于校验登录参数。
func Authenticator(ctx context.Context) (interface{}, error) {
	var (
		r   = g.RequestFromCtx(ctx)
		req *v1.LoginReq
	)
	//从 HTTP 请求中读取用户传来的参数，并绑定到 Go 的结构体里
	if err := r.Parse(&req); err != nil {
		return nil, err
	}

	eUser, err := User().Get(ctx, &do.User{Username: req.Username})
	if err != nil {
		return nil, err
	}

	if eUser == nil {
		return nil, gerror.New("用户不存在")
	}

	if !util.ComparePassword(eUser.Password, req.Password) {
		return nil, gerror.New("密码错误")
	}

	if !eUser.Enabled {
		return nil, gerror.New("该用户处于禁用状态")
	}

	return eUser, nil
}
