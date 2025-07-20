package role

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"omp/api/role/v1"
)

func (c *ControllerV1) GetPageLst(ctx context.Context, req *v1.GetPageLstReq) (res *v1.GetPageLstRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
