package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"gtodolist/app/user/cmd/api/internal/svc"
	"gtodolist/app/user/cmd/api/internal/types"
	"gtodolist/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// 向 rpc 发送请求
	registerResp, err := l.svcCtx.UserRpcClient.Register(l.ctx, &pb.RegisterReq{
		Username: req.Username,
		Password: req.Password,
	})

	resp = &types.RegisterResp{}
	_ = copier.Copy(resp, registerResp)
	return resp, err
}
