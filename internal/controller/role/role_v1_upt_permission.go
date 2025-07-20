package role

import (
	"context"
	"omp/internal/model/do"
	"omp/internal/service"

	"omp/api/role/v1"
)

func (c *ControllerV1) UptPermission(ctx context.Context, req *v1.UptPermissionReq) (res *v1.UptPermissionRes, err error) {
	err = service.Role().Upt(ctx, &do.Role{Id: req.Id, Permission: req.PermissionIds})
	return
}
