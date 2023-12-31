// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"gtodolist/app/user/cmd/rpc/internal/logic"
	"gtodolist/app/user/cmd/rpc/internal/svc"
	"gtodolist/app/user/cmd/rpc/pb"
)

type UserrpcServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUserrpcServer
}

func NewUserrpcServer(svcCtx *svc.ServiceContext) *UserrpcServer {
	return &UserrpcServer{
		svcCtx: svcCtx,
	}
}

func (s *UserrpcServer) Register(ctx context.Context, in *pb.RegisterReq) (*pb.RegisterResp, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

func (s *UserrpcServer) Login(ctx context.Context, in *pb.LoginReq) (*pb.LoginResp, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UserrpcServer) GenerateToken(ctx context.Context, in *pb.GenerateTokenReq) (*pb.GenerateTokenResp, error) {
	l := logic.NewGenerateTokenLogic(ctx, s.svcCtx)
	return l.GenerateToken(in)
}
