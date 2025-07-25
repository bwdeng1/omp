// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"omp/api/user/v1"
)

type IUserV1 interface {
	Add(ctx context.Context, req *v1.AddReq) (res *v1.AddRes, err error)
	GetLst(ctx context.Context, req *v1.GetLstReq) (res *v1.GetLstRes, err error)
	GetPageLst(ctx context.Context, req *v1.GetPageLstReq) (res *v1.GetPageLstRes, err error)
	Upt(ctx context.Context, req *v1.UptReq) (res *v1.UptRes, err error)
	UptPassword(ctx context.Context, req *v1.UptPasswordReq) (res *v1.UptPasswordRes, err error)
	UptEnabled(ctx context.Context, req *v1.UptEnabledReq) (res *v1.UptEnabledRes, err error)
	Del(ctx context.Context, req *v1.DelReq) (res *v1.DelRes, err error)
}
