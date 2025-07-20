package user

import (
	"context"
	"omp/internal/model/do"
	"omp/internal/service"

	"omp/api/user/v1"
)

func (c *ControllerV1) Del(ctx context.Context, req *v1.DelReq) (res *v1.DelRes, err error) {
	err = service.User().Del(ctx, &do.User{Id: req.Id})
	return
}
