package user

import (
	"context"
	"omp/internal/model/do"
	"omp/internal/service"

	"omp/api/user/v1"
)

func (c *ControllerV1) UptPassword(ctx context.Context, req *v1.UptPasswordReq) (res *v1.UptPasswordRes, err error) {
	err = service.User().Upt(ctx, &do.User{Id: req.Id, Password: req.Password})
	return
}
