package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"omp/api"
	"omp/internal/model/entity"
	"omp/internal/model/mid"
)

// AddReq 新增用户的请求结构体
type AddReq struct {
	g.Meta    `method:"post" path:"/user" summary:"新增用户" tags:"用户"`
	*mid.User // 嵌套了 mid.User，复用了字段和校验规则
}

// 新增成功无需返回特定数据，响应一个空对象即可
type AddRes struct{}

type GetLstReq struct {
	g.Meta `method:"get" path:"/user/list" summary:"获取用户列表" tags:"用户"`
}

type GetLstRes struct {
	List []*entity.User `json:"list"`
}

// GetPageLstReq 获取用户分页列表的请求结构体
type GetPageLstReq struct {
	g.Meta          `method:"get" path:"/user/page-list" summary:"分页获取用户列表" tags:"用户"`
	*api.PageLstReq // 嵌套了通用的分页请求参数
}

// GetPageLstRes 获取用户分页列表的响应结构体
type GetPageLstRes struct {
	*api.PageLstRes[*entity.User] // 嵌套了通用的分页响应结构，并指定列表项类型为 entity.User
}

type UptReq struct {
	g.Meta `method:"put" path:"/user/{id}" summary:"更新用户" tags:"用户"`
	Id     int ` v:"min:1#id必须" path:"id"`
	*mid.User
}

type UptRes struct{}

type UptPasswordReq struct {
	g.Meta   `method:"patch" path:"/user/{id}/password" summary:"更新用户密码" tags:"用户"`
	Id       int    ` v:"min:1#id必须" path:"id"`
	Password string `v:"required|length:6,30#请输入密码|密码长度为:{min}到:{max}位"`
}

type UptPasswordRes struct{}

type UptEnabledReq struct {
	g.Meta  `method:"patch" path:"/user/{id}/enabled" summary:"更新用户密码" tags:"用户"`
	Id      int  `v:"min:1#id必须" path:"id"`
	Enabled bool `v:"required" json:"enabled"`
}

type UptEnabledRes struct{}

type DelReq struct {
	g.Meta `method:"delete" path:"/user/{id}" summary:"删除用户" tags:"用户"`
	Id     int ` v:"min:1#id必须" path:"id"`
}

type DelRes struct{}
