package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"gtodolist/app/task/cmd/api/internal/config"
	"gtodolist/app/task/cmd/rpc/taskrpc"
)

type ServiceContext struct {
	Config        config.Config
	TaskRpcClient taskrpc.Taskrpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		TaskRpcClient: taskrpc.NewTaskrpc(zrpc.MustNewClient(c.TaskRpcConfig)),
	}
}
