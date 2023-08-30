package svc

import (
	"github.com/SpectatorNan/gorm-zero/gormc"
	"gtodolist/app/user/cmd/rpc/internal/config"
	"gtodolist/app/user/model"
	"log"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gormc.ConnectMysql(c.Mysql)
	if err != nil {
		log.Fatal(err)
	}
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(db, c.Cache),
	}
}
