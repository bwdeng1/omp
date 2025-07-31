package nginx

import (
	"bytes"
	"context"
	"fmt"
	"omp/internal/dao"
	"omp/internal/model/entity"
	"os/exec"
	"strings"
	"sync"

	"github.com/gogf/gf/v2/frame/g"

	"omp/api/nginx/v1"
	"omp/internal/service"
)

type sNginx struct{}

func init() {
	service.RegisterNginx(New())
}

func New() *sNginx {
	return &sNginx{}
}

func (s *sNginx) SwitchBackend(ctx context.Context, bizKey, target string) ([]v1.SwitchResult, error) {
	instanceIds := g.Cfg().MustGet(ctx, "nginx.instanceIds").Strings()
	scriptPathKey := fmt.Sprintf("nginx.scripts.%s", bizKey)
	scriptPath := g.Cfg().MustGet(ctx, scriptPathKey).String()
	if scriptPath == "" {
		return nil, fmt.Errorf("配置缺失: 未在config.yaml中找到业务线 '%s' 对应的脚本路径", bizKey)
	}
	regionId := g.Cfg().MustGet(ctx, "aliyun.regionId").String()

	var wg sync.WaitGroup
	resultsChan := make(chan v1.SwitchResult, len(instanceIds))

	for _, instanceId := range instanceIds {
		wg.Add(1)
		go func(id string) {
			defer wg.Done()
			commandToRun := fmt.Sprintf("sudo %s %s", scriptPath, target)
			result := s.executeSwitchOnInstanceCLI(ctx, regionId, id, commandToRun)
			resultsChan <- result
		}(instanceId)
	}

	wg.Wait()
	close(resultsChan)

	var finalResults []v1.SwitchResult
	for result := range resultsChan {
		finalResults = append(finalResults, result)
	}

	return finalResults, nil
}

// executeSwitchOnInstanceCLI 使用命令行调用 aliyun ecs RunCommand
func (s *sNginx) executeSwitchOnInstanceCLI(ctx context.Context, regionId, instanceID, command string) v1.SwitchResult {

	args := []string{
		"ecs",
		"RunCommand",
		"--RegionId", regionId,
		"--InstanceId.1", instanceID,
		"--Type", "RunShellScript",
		"--CommandContent", command,
	}

	cmd := exec.Command("aliyun", args...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	g.Log().Infof(ctx, "即将执行CLI命令: aliyun %s", strings.Join(args, " "))

	err := cmd.Run()

	if err != nil {
		g.Log().Errorf(ctx, "执行aliyun CLI失败, InstanceID: %s, Stderr: %s", instanceID, stderr.String())
		return v1.SwitchResult{InstanceID: instanceID, Status: "error", Output: stderr.String()}
	}

	return v1.SwitchResult{
		InstanceID: instanceID,
		InvokeID:   "executed_by_cli",
		Status:     "success",
		Output:     "通过 aliyun CLI 命令已成功下发。",
	}
}

// UpdateStatus 更新或插入Nginx切换状态
func (s *sNginx) UpdateStatus(ctx context.Context, bizKey string, activeEnv int) error {
	// 使用 Save 方法，如果记录已存在则更新，不存在则插入，非常方便
	_, err := dao.NginxSwitchStatus.Ctx(ctx).Data(g.Map{
		"biz_key":    bizKey,
		"active_env": activeEnv,
		// 您可以从 ctx 中获取当前操作的用户信息，并更新 LastOperator 字段
		// "last_operator": service.Session().GetUser(ctx).Username,
	}).Save()

	if err != nil {
		g.Log().Errorf(ctx, "更新Nginx状态失败, BizKey: %s, Error: %v", bizKey, err)
		return err
	}

	g.Log().Infof(ctx, "Nginx状态更新成功, BizKey: %s, ActiveEnv: %d", bizKey, activeEnv)
	return nil
}

// GetStatus 查询Nginx切换状态
func (s *sNginx) GetStatus(ctx context.Context, bizKey string) (*entity.NginxSwitchStatus, error) {
	var status *entity.NginxSwitchStatus

	err := dao.NginxSwitchStatus.Ctx(ctx).Where(dao.NginxSwitchStatus.Columns().BizKey, bizKey).Scan(&status)
	if err != nil {
		g.Log().Errorf(ctx, "查询Nginx状态失败, BizKey: %s, Error: %v", bizKey, err)
		return nil, err
	}

	// 如果查询不到记录，可以返回一个默认值，例如环境0
	if status == nil {
		return &entity.NginxSwitchStatus{
			BizKey:    bizKey,
			ActiveEnv: 0, // 0 代表未知或双活
		}, nil
	}

	return status, nil
}
