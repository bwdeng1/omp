// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
)

// Permission is the golang structure for table permission.
type Permission struct {
	Id         int         `json:"id"         orm:"id"          description:""`                  //
	Title      string      `json:"title"      orm:"title"       description:"标题"`                // 标题
	Name       string      `json:"name"       orm:"name"        description:"路由名称"`              // 路由名称
	Type       int         `json:"type"       orm:"type"        description:"类型:1-目录,2-菜单,3-功能"` // 类型:1-目录,2-菜单,3-功能
	FRoute     string      `json:"fRoute"     orm:"f_route"     description:"前端路由路径"`            // 前端路由路径
	BRoutes    *gjson.Json `json:"bRoutes"    orm:"b_routes"    description:"后端路由路径"`            // 后端路由路径
	Redirect   string      `json:"redirect"   orm:"redirect"    description:"重定向路径"`             // 重定向路径
	Icon       string      `json:"icon"       orm:"icon"        description:"图标"`                // 图标
	Rank       int         `json:"rank"       orm:"rank"        description:"排序"`                // 排序
	ShowLink   bool        `json:"showLink"   orm:"show_link"   description:"是否在菜单中展示"`          // 是否在菜单中展示
	ShowParent bool        `json:"showParent" orm:"show_parent" description:"是否展示父级菜单"`          // 是否展示父级菜单
	KeepAlive  bool        `json:"keepAlive"  orm:"keep_alive"  description:"页面缓存"`              // 页面缓存
	ParentId   int         `json:"parentId"   orm:"parent_id"   description:"父级权限 id"`           // 父级权限 id
}
