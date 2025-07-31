// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NginxSwitchStatusDao is the data access object for the table nginx_switch_status.
type NginxSwitchStatusDao struct {
	table    string                   // table is the underlying table name of the DAO.
	group    string                   // group is the database configuration group name of the current DAO.
	columns  NginxSwitchStatusColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler       // handlers for customized model modification.
}

// NginxSwitchStatusColumns defines and stores column names for the table nginx_switch_status.
type NginxSwitchStatusColumns struct {
	Id           string //
	BizKey       string // 业务线标识, 例如 mealpal1
	ActiveEnv    string // 当前生效的环境: 1代表第一套, 2代表第二套, 0代表双活
	UpdatedAt    string // 最后更新时间
	LastOperator string // 最后操作人
}

// nginxSwitchStatusColumns holds the columns for the table nginx_switch_status.
var nginxSwitchStatusColumns = NginxSwitchStatusColumns{
	Id:           "id",
	BizKey:       "biz_key",
	ActiveEnv:    "active_env",
	UpdatedAt:    "updated_at",
	LastOperator: "last_operator",
}

// NewNginxSwitchStatusDao creates and returns a new DAO object for table data access.
func NewNginxSwitchStatusDao(handlers ...gdb.ModelHandler) *NginxSwitchStatusDao {
	return &NginxSwitchStatusDao{
		group:    "default",
		table:    "nginx_switch_status",
		columns:  nginxSwitchStatusColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *NginxSwitchStatusDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *NginxSwitchStatusDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *NginxSwitchStatusDao) Columns() NginxSwitchStatusColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *NginxSwitchStatusDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *NginxSwitchStatusDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *NginxSwitchStatusDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
