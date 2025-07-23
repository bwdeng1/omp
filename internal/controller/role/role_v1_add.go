package role

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"omp/internal/model/entity"
	"omp/internal/service"

	"omp/api/role/v1"
)

func (c *ControllerV1) Add(ctx context.Context, req *v1.AddReq) (res *v1.AddRes, err error) {
	in := new(entity.Role)
	if err = gconv.Struct(req, in); err != nil {
		return
	}

	err = service.Role().Add(ctx, in)
	return
}
