package logic

import (
	"context"

	"gtodolist/app/task/cmd/api/internal/svc"
	"gtodolist/app/task/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Task_searchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTask_searchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Task_searchLogic {
	return &Task_searchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Task_searchLogic) Task_search(req *types.SearchReq) (resp *types.SearchResp, err error) {
	// todo: add your logic here and delete this line

	return
}
