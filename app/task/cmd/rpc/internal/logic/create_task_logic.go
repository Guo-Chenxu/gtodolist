package logic

import (
	"context"

	"gtodolist/app/task/cmd/rpc/internal/svc"
	"gtodolist/app/task/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTaskLogic {
	return &CreateTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateTaskLogic) CreateTask(in *pb.CreateReq) (*pb.CreateReq, error) {
	// todo: add your logic here and delete this line

	return &pb.CreateReq{}, nil
}
