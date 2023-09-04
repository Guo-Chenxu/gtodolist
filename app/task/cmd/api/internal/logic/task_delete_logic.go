package logic

import (
	"context"

	"gtodolist/app/task/cmd/api/internal/svc"
	"gtodolist/app/task/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Task_deleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTask_deleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Task_deleteLogic {
	return &Task_deleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Task_deleteLogic) Task_delete(req *types.DeleteReq) (resp *types.DeleteResp, err error) {
	// todo: add your logic here and delete this line

	return
}
