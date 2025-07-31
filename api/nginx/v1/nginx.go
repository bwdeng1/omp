package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SwitchReq 定义了后端切换的API请求结构体
type SwitchReq struct {
	g.Meta `path:"/nginx/switch" method:"post" tags:"Nginx切换" summary:"执行Nginx后端服务切换"`
	// 业务线标识，用于在配置中查找对应的脚本。
	BizKey string `json:"bizKey" v:"required#业务线标识不能为空" dc:"业务线标识 (如: mealpal1)"`
	// 切换目标，这个是您脚本需要接收的参数，例如 "blue" 或 "green"
	Target string `json:"target" v:"required#切换目标不能为空" dc:"切换目标 (传给脚本的参数)"`
}

// 单台服务器的执行结果
type SwitchResult struct {
	InstanceID string `json:"instanceId" dc:"ECS实例ID"`
	InvokeID   string `json:"invokeId"   dc:"云助手执行ID"`
	Status     string `json:"status"     dc:"执行状态 (success/error)"`
	Output     string `json:"output"     dc:"执行输出或错误信息"`
}

// SwitchRes 定义了后端切换的API响应结构体
type SwitchRes struct {
	Results []SwitchResult `json:"results" dc:"各台服务器的执行结果"`
}

// UpdateStatusReq 定义了更新Nginx后端状态的API请求结构体
type UpdateStatusReq struct {
	g.Meta    `path:"/nginx/status" method:"post" tags:"Nginx切换" summary:"更新Nginx后端状态"`
	BizKey    string `json:"bizKey" v:"required#业务线标识不能为空"`
	ActiveEnv int    `json:"activeEnv" v:"required|in:0,1,2#状态值必须是0,1,2之一"`
}

// UpdateStatusRes 定义了更新状态的API响应
type UpdateStatusRes struct{}

// GetStatusReq 定义了查询Nginx后端状态的API请求结构体
type GetStatusReq struct {
	g.Meta `path:"/nginx/status" method:"get" tags:"Nginx切换" summary:"查询Nginx后端状态"`
	BizKey string `json:"bizKey" v:"required#业务线标识不能为空"`
}

// GetStatusRes 定义了查询状态的API响应
type GetStatusRes struct {
	BizKey    string `json:"bizKey"`
	ActiveEnv int    `json:"activeEnv"`
}
