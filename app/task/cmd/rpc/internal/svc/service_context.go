package svc

import (
	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gtodolist/app/task/cmd/rpc/internal/config"
	"gtodolist/app/task/model"
	"log"
)

type ServiceContext struct {
	Config      config.Config
	TaskModel   model.TaskModel
	RedisClient *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gormc.ConnectMysql(c.Mysql)
	if err != nil {
		log.Fatal(err)
	}
	return &ServiceContext{
		Config:      c,
		TaskModel:   model.NewTaskModel(db, c.Cache),
		RedisClient: redis.MustNewRedis(c.Redis.RedisConf),
	}
}
