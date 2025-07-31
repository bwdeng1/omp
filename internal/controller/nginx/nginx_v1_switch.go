package nginx

import (
	"context"
	"omp/internal/service"

	"omp/api/nginx/v1"
)

func (c *ControllerV1) Switch(ctx context.Context, req *v1.SwitchReq) (res *v1.SwitchRes, err error) {
	// 调用服务时，传入新的 bizKey 参数
	results, err := service.Nginx().SwitchBackend(ctx, req.BizKey, req.Target)
	if err != nil {
		return nil, err
	}

	res = &v1.SwitchRes{
		Results: results,
	}

	return res, nil
}
