package nginx

import (
	"context"
	"omp/internal/service"

	"omp/api/nginx/v1"
)

func (c *ControllerV1) GetStatus(ctx context.Context, req *v1.GetStatusReq) (res *v1.GetStatusRes, err error) {
	status, err := service.Nginx().GetStatus(ctx, req.BizKey)
	if err != nil {
		return nil, err
	}

	res = &v1.GetStatusRes{
		BizKey:    status.BizKey,
		ActiveEnv: status.ActiveEnv,
	}
	return
}
