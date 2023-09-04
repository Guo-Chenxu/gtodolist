package logic

import (
	"context"

	"gtodolist/app/task/cmd/api/internal/svc"
	"gtodolist/app/task/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Task_listLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTask_listLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Task_listLogic {
	return &Task_listLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Task_listLogic) Task_list(req *types.ListReq) (resp *types.ListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
