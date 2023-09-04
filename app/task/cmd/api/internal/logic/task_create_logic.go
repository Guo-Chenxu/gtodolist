package logic

import (
	"context"

	"gtodolist/app/task/cmd/api/internal/svc"
	"gtodolist/app/task/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Task_createLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTask_createLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Task_createLogic {
	return &Task_createLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Task_createLogic) Task_create(req *types.CreateReq) (resp *types.CreateResp, err error) {
	// todo: add your logic here and delete this line

	return
}
