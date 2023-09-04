package svc

import (
	"github.com/SpectatorNan/gorm-zero/gormc"
	"gtodolist/app/task/cmd/rpc/internal/config"
	"gtodolist/app/task/model"
	"log"
)

type ServiceContext struct {
	Config    config.Config
	TaskModel model.TaskModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gormc.ConnectMysql(c.Mysql)
	if err != nil {
		log.Fatal(err)
	}
	return &ServiceContext{
		Config:    c,
		TaskModel: model.NewTaskModel(db, c.Cache),
	}
}
