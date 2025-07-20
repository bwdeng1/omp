package permission

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"omp/internal/model/entity"
	"omp/internal/service"

	"omp/api/permission/v1"
)

func (c *ControllerV1) Add(ctx context.Context, req *v1.AddReq) (res *v1.AddRes, err error) {
	ePermission := new(entity.Permission)
	if err = gconv.Struct(req, ePermission); err != nil {
		return
	}
	err = service.Permission().Add(ctx, ePermission)
	return
}
