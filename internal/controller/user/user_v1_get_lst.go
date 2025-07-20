package user

import (
	"context"
	"omp/api/user/v1"
	"omp/internal/service"
)

func (c *ControllerV1) GetLst(ctx context.Context, req *v1.GetLstReq) (res *v1.GetLstRes, err error) {
	res = new(v1.GetLstRes)
	res.List, err = service.User().GetLst(ctx)
	return
}
