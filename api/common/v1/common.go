package v1

import "github.com/gogf/gf/v2/frame/g"

// GetGitBranchLstReq 定义了获取 Git 分支列表的API请求结构体
type GetGitBranchLstReq struct {
	g.Meta `method:"get" path:"/common/git-branch-list" summary:"获取指定 Git 仓库的分支名称列表" tags:"通用"`
	// Git 仓库的克隆地址
	GitUrl   string `v:"required" p:"gitUrl"`
	SecretId int    `v:"required" p:"secretId"`
}

// GetGitBranchLstRes 定义了获取 Git 分支列表的API响应结构体
type GetGitBranchLstRes struct {
	BranchLst []string `json:"branches"`
}
