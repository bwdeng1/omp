// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package nginx

import (
	"context"

	"omp/api/nginx/v1"
)

type INginxV1 interface {
	Switch(ctx context.Context, req *v1.SwitchReq) (res *v1.SwitchRes, err error)
	UpdateStatus(ctx context.Context, req *v1.UpdateStatusReq) (res *v1.UpdateStatusRes, err error)
	GetStatus(ctx context.Context, req *v1.GetStatusReq) (res *v1.GetStatusRes, err error)
}
