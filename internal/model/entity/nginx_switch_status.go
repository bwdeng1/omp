// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// NginxSwitchStatus is the golang structure for table nginx_switch_status.
type NginxSwitchStatus struct {
	Id           int         `json:"id"           orm:"id"            description:""`                               //
	BizKey       string      `json:"bizKey"       orm:"biz_key"       description:"业务线标识, 例如 mealpal1"`             // 业务线标识, 例如 mealpal1
	ActiveEnv    int         `json:"activeEnv"    orm:"active_env"    description:"当前生效的环境: 1代表第一套, 2代表第二套, 0代表双活"` // 当前生效的环境: 1代表第一套, 2代表第二套, 0代表双活
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:"最后更新时间"`                         // 最后更新时间
	LastOperator string      `json:"lastOperator" orm:"last_operator" description:"最后操作人"`                          // 最后操作人
}
