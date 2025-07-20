package common

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"omp/api/common/v1"
)

func (c *ControllerV1) GetGitBranchLst(ctx context.Context, req *v1.GetGitBranchLstReq) (res *v1.GetGitBranchLstRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
