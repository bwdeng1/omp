// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "omp/api/nginx/v1"
	"omp/internal/model/entity"
)

type (
	INginx interface {
		SwitchBackend(ctx context.Context, bizKey string, target string) ([]v1.SwitchResult, error)
		// UpdateStatus 更新或插入Nginx切换状态
		UpdateStatus(ctx context.Context, bizKey string, activeEnv int) error
		// GetStatus 查询Nginx切换状态
		GetStatus(ctx context.Context, bizKey string) (*entity.NginxSwitchStatus, error)
	}
)

var (
	localNginx INginx
)

func Nginx() INginx {
	if localNginx == nil {
		panic("implement not found for interface INginx, forgot register?")
	}
	return localNginx
}

func RegisterNginx(i INginx) {
	localNginx = i
}
