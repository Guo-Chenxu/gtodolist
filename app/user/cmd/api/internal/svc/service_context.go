package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"gtodolist/app/user/cmd/api/internal/config"
	"gtodolist/app/user/cmd/rpc/userrpc"
)

type ServiceContext struct {
	Config          config.Config
	LoginMiddleware rest.Middleware
	UserRpcClient   userrpc.Userrpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserRpcClient: userrpc.NewUserrpc(zrpc.MustNewClient(c.UserRpcConfig)),
	}
}
