package logic

import (
	"context"

	"gtodolist/app/task/cmd/api/internal/svc"
	"gtodolist/app/task/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskSearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTaskSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskSearchLogic {
	return &TaskSearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TaskSearchLogic) TaskSearch(req *types.SearchReq) (resp *types.SearchResp, err error) {
	// todo: add your logic here and delete this line

	return
}
