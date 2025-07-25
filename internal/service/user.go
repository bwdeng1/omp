// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"omp/api"
	"omp/internal/model/comb"
	"omp/internal/model/do"
	"omp/internal/model/entity"
)

type (
	IUser interface {
		Add(ctx context.Context, in *entity.User) (err error)
		Upt(ctx context.Context, in *do.User) (err error)
		GetLst(ctx context.Context) (out []*entity.User, err error)
		GetPageLst(ctx context.Context, in *api.PageLstReq) (out *api.PageLstRes[*entity.User], err error)
		Get(ctx context.Context, userDo *do.User) (out *entity.User, err error)
		GetComb(ctx context.Context, userDo *do.User) (out *comb.User, err error)
		GetCombLst(ctx context.Context) (out []*comb.User, err error)
		Del(ctx context.Context, in *do.User) (err error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
