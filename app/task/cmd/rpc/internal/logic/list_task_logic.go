package logic

import (
	"context"

	"gtodolist/app/task/cmd/rpc/internal/svc"
	"gtodolist/app/task/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListTaskLogic {
	return &ListTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListTaskLogic) ListTask(in *pb.ListReq) (*pb.ListResp, error) {
	// todo: add your logic here and delete this line

	return &pb.ListResp{}, nil
}
