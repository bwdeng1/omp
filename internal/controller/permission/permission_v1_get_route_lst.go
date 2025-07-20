package permission

import (
	"context"
	"omp/internal/service"

	"omp/api/permission/v1"
)

func (c *ControllerV1) GetRouteLst(ctx context.Context, req *v1.GetRouteLstReq) (res *v1.GetRouteLstRes, err error) {
	res = new(v1.GetRouteLstRes)
	res.List, err = service.Permission().GetRouteLst(ctx)
	return
}
