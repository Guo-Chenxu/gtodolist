package logic

import (
	"context"

	"gtodolist/app/task/cmd/rpc/internal/svc"
	"gtodolist/app/task/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShowTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewShowTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShowTaskLogic {
	return &ShowTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ShowTaskLogic) ShowTask(in *pb.ShowReq) (*pb.ShowResp, error) {
	// todo: add your logic here and delete this line

	return &pb.ShowResp{}, nil
}
