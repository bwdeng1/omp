// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id        int         `json:"id"        orm:"id"         description:""`        //
	Username  string      `json:"username"  orm:"username"   description:"用户名"`     // 用户名
	Password  string      `json:"password"  orm:"password"   description:"密码"`      // 密码
	Phone     string      `json:"phone"     orm:"phone"      description:"手机号码"`    // 手机号码
	Email     string      `json:"email"     orm:"email"      description:"邮箱"`      // 邮箱
	RealName  string      `json:"realName"  orm:"real_name"  description:"真实姓名"`    // 真实姓名
	Enabled   bool        `json:"enabled"   orm:"enabled"    description:"是否启用状态"`  // 是否启用状态
	RoleIds   *gjson.Json `json:"roleIds"   orm:"role_ids"   description:"角色 id"`   // 角色 id
	DeptId    int         `json:"deptId"    orm:"dept_id"    description:"所属部门 id"` // 所属部门 id
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`    // 更新时间
}
