package public

import (
	"context"
	"omp/internal/model/do"
	"omp/internal/service"
	"time"

	"omp/api/public/v1"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	token, expires := service.Auth().LoginHandler(ctx)
	refreshAfter := time.Now().Add(expires.Sub(time.Now()) / 2).UnixMilli()
	eUser, err := service.User().GetComb(ctx, &do.User{Username: req.Username})
	if err != nil {
		return nil, err
	}
	res = &v1.LoginRes{
		Username:     eUser.Username,
		RealName:     eUser.RealName,
		Token:        token,
		Expires:      expires.UnixMilli(),
		RefreshAfter: refreshAfter,
		Roles:        eUser.RoleCodes(),
	}
	return
}
