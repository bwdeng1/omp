package permission

import (
	"context"
	"omp/internal/model/do"
	"omp/internal/service"

	"omp/api/permission/v1"
)

func (c *ControllerV1) UptShowLink(ctx context.Context, req *v1.UptShowLinkReq) (res *v1.UptShowLinkRes, err error) {
	err = service.Permission().Upt(ctx, &do.Permission{Id: req.Id, ShowLink: req.Enabled})
	return
}
