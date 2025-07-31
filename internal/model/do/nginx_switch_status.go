// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// NginxSwitchStatus is the golang structure of table nginx_switch_status for DAO operations like Where/Data.
type NginxSwitchStatus struct {
	g.Meta       `orm:"table:nginx_switch_status, do:true"`
	Id           interface{} //
	BizKey       interface{} // 业务线标识, 例如 mealpal1
	ActiveEnv    interface{} // 当前生效的环境: 1代表第一套, 2代表第二套, 0代表双活
	UpdatedAt    *gtime.Time // 最后更新时间
	LastOperator interface{} // 最后操作人
}
