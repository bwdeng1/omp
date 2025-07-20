package user

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"omp/api/user/v1"
	"omp/internal/model/do"
	"omp/internal/service"
)

func (c *ControllerV1) Upt(ctx context.Context, req *v1.UptReq) (res *v1.UptRes, err error) {
	uptDo := new(do.User)
	if err = gconv.Struct(req, uptDo); err != nil {
		return
	}
	err = service.User().Upt(ctx, uptDo)
	return

}
