package user

import (
	"context"
	"omp/internal/model/do"
	"omp/internal/service"

	"omp/api/user/v1"
)

func (c *ControllerV1) UptEnabled(ctx context.Context, req *v1.UptEnabledReq) (res *v1.UptEnabledRes, err error) {
	err = service.User().Upt(ctx, &do.User{Id: req.Id, Enabled: req.Enabled})
	return
}
