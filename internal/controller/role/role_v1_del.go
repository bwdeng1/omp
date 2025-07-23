package role

import (
	"context"
	"omp/internal/model/do"
	"omp/internal/service"

	"omp/api/role/v1"
)

func (c *ControllerV1) Del(ctx context.Context, req *v1.DelReq) (res *v1.DelRes, err error) {
	err = service.Role().Del(ctx, &do.Role{Id: req.Id})
	return
}
