package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"gtodolist/app/user/cmd/api/internal/config"
	"gtodolist/app/user/cmd/api/internal/middleware"
)

type ServiceContext struct {
	Config          config.Config
	LoginMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		LoginMiddleware: middleware.NewLoginMiddleware().Handle,
	}
}
