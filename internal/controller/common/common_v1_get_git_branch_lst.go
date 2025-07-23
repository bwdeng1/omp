package common

import (
	"context"
	"omp/internal/service"

	"omp/api/common/v1"
)

func (c *ControllerV1) GetGitBranchLst(ctx context.Context, req *v1.GetGitBranchLstReq) (res *v1.GetGitBranchLstRes, err error) {
	res = new(v1.GetGitBranchLstRes)
	res.BranchLst, err = service.Common().GetGitBranchLst(ctx, req.GitUrl, req.SecretId)
	return
}
