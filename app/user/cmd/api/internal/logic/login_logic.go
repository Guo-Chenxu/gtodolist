package logic

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"gtodolist/app/user/cmd/rpc/pb"

	"gtodolist/app/user/cmd/api/internal/svc"
	"gtodolist/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	loginResp, err := l.svcCtx.UserRpcClient.Login(l.ctx, &pb.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})

	fmt.Println(loginResp)
	if err != nil {
		return &types.LoginResp{
			Status:  int(loginResp.Status),
			Message: loginResp.Message,
			Error:   err.Error(),
		}, err
	}

	resp = &types.LoginResp{}
	_ = copier.Copy(resp, loginResp)
	return resp, err
}
