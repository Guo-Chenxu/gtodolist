package logic

import (
	"context"

	"gtodolist/app/task/cmd/rpc/internal/svc"
	"gtodolist/app/task/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTaskLogic {
	return &DeleteTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteTaskLogic) DeleteTask(in *pb.DeleteReq) (*pb.DeleteResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DeleteResp{}, nil
}
