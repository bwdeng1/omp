package user

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"omp/api/user/v1"
	"omp/internal/model/entity"
	"omp/internal/service"
)

func (c *ControllerV1) Add(ctx context.Context, req *v1.AddReq) (res *v1.AddRes, err error) {
	eUser := new(entity.User)
	//gconv.Struct 将 req (API层结构体) 转换为 eUser (Logic/DAO层使用的实体)。
	if err = gconv.Struct(req, eUser); err != nil {
		return
	}

	eUser.Password = "517na.com"
	err = service.User().Add(ctx, eUser)
	return
}
