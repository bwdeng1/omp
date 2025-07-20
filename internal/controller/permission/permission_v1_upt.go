package permission

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"omp/internal/model/do"
	"omp/internal/service"

	"omp/api/permission/v1"
)

func (c *ControllerV1) Upt(ctx context.Context, req *v1.UptReq) (res *v1.UptRes, err error) {
	uptDo := new(do.Permission)
	if err = gconv.Struct(req, uptDo); err != nil {
		return
	}
	err = service.Permission().Upt(ctx, uptDo)
	return
}
