package permission

import (
	"context"
	"omp/internal/model/do"
	"omp/internal/service"

	"omp/api/permission/v1"
)

func (c *ControllerV1) Del(ctx context.Context, req *v1.DelReq) (res *v1.DelRes, err error) {
	err = service.Permission().Del(ctx, &do.Permission{Id: req.Id})
	return
}
