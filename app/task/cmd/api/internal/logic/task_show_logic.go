package logic

import (
	"context"

	"gtodolist/app/task/cmd/api/internal/svc"
	"gtodolist/app/task/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Task_showLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTask_showLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Task_showLogic {
	return &Task_showLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Task_showLogic) Task_show(req *types.ShowReq) (resp *types.ShowResp, err error) {
	// todo: add your logic here and delete this line

	return
}
