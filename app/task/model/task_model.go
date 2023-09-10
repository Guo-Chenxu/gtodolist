package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
	"strconv"
)

var _ TaskModel = (*customTaskModel)(nil)

type (
	// TaskModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTaskModel.
	TaskModel interface {
		taskModel
		customTaskLogicModel
	}

	customTaskModel struct {
		*defaultTaskModel
	}

	customTaskLogicModel interface {
		myTaskModel
	}
)

// NewTaskModel returns a model for the database table.
func NewTaskModel(conn *gorm.DB, c cache.CacheConf) TaskModel {
	return &customTaskModel{
		defaultTaskModel: newTaskModel(conn, c),
	}
}

func (m *defaultTaskModel) customCacheKeys(data *Task) []string {
	if data == nil {
		return []string{}
	}

	// todo: 这里会缓存不一致, 无法删除旧缓存
	return []string{
		cacheGtodolistTaskListPrefix + strconv.FormatInt(data.Uid, 10),
	}
}
