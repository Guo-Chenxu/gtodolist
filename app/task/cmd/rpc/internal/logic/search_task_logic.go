package logic

import (
	"context"

	"gtodolist/app/task/cmd/rpc/internal/svc"
	"gtodolist/app/task/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchTaskLogic {
	return &SearchTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchTaskLogic) SearchTask(in *pb.SearchReq) (*pb.SearchResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchResp{}, nil
}
