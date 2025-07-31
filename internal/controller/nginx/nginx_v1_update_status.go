package nginx

import (
	"context"
	"omp/internal/service"

	"omp/api/nginx/v1"
)

func (c *ControllerV1) UpdateStatus(ctx context.Context, req *v1.UpdateStatusReq) (res *v1.UpdateStatusRes, err error) {
	err = service.Nginx().UpdateStatus(ctx, req.BizKey, req.ActiveEnv)
	return
}
