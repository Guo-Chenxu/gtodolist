package logic

import (
	"context"

	"gtodolist/app/task/cmd/api/internal/svc"
	"gtodolist/app/task/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskShowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTaskShowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskShowLogic {
	return &TaskShowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TaskShowLogic) TaskShow(req *types.ShowReq) (resp *types.ShowResp, err error) {
	// todo: add your logic here and delete this line

	return
}
