// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PermissionDao is the data access object for the table permission.
type PermissionDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  PermissionColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// PermissionColumns defines and stores column names for the table permission.
type PermissionColumns struct {
	Id         string //
	Title      string // 标题
	Name       string // 路由名称
	Type       string // 类型:1-目录,2-菜单,3-功能
	FRoute     string // 前端路由路径
	BRoutes    string // 后端路由路径
	Redirect   string // 重定向路径
	Icon       string // 图标
	Rank       string // 排序
	ShowLink   string // 是否在菜单中展示
	ShowParent string // 是否展示父级菜单
	KeepAlive  string // 页面缓存
	ParentId   string // 父级权限 id
}

// permissionColumns holds the columns for the table permission.
var permissionColumns = PermissionColumns{
	Id:         "id",
	Title:      "title",
	Name:       "name",
	Type:       "type",
	FRoute:     "f_route",
	BRoutes:    "b_routes",
	Redirect:   "redirect",
	Icon:       "icon",
	Rank:       "rank",
	ShowLink:   "show_link",
	ShowParent: "show_parent",
	KeepAlive:  "keep_alive",
	ParentId:   "parent_id",
}

// NewPermissionDao creates and returns a new DAO object for table data access.
func NewPermissionDao(handlers ...gdb.ModelHandler) *PermissionDao {
	return &PermissionDao{
		group:    "default",
		table:    "permission",
		columns:  permissionColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PermissionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PermissionDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PermissionDao) Columns() PermissionColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PermissionDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PermissionDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PermissionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
