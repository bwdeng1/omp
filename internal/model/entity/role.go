// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// Role is the golang structure for table role.
type Role struct {
	Id         int         `json:"id"         orm:"id"         description:""`     //
	Name       string      `json:"name"       orm:"name"       description:"角色名称"` // 角色名称
	Code       string      `json:"code"       orm:"code"       description:"角色代码"` // 角色代码
	Permission *gjson.Json `json:"permission" orm:"permission" description:"关联权限"` // 关联权限
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at" description:"更新时间"` // 更新时间
}
