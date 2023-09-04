package logic

import (
	"context"

	"gtodolist/app/task/cmd/api/internal/svc"
	"gtodolist/app/task/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Task_updateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTask_updateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Task_updateLogic {
	return &Task_updateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Task_updateLogic) Task_update(req *types.UpdateReq) (resp *types.UpdateResp, err error) {
	// todo: add your logic here and delete this line

	return
}
