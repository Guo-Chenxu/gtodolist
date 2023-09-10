package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

var (
	cacheGtodolistTaskListPrefix = "cache:gtodolist:task:list:"
)

type (
	myTaskModel interface {
		List(ctx context.Context, uid, page, pageSize int64) ([]Task, error)
	}
)

func (m *defaultTaskModel) List(ctx context.Context, uid, page, pageSize int64) ([]Task, error) {
	key := fmt.Sprintf("%s%v:%v:%v", cacheGtodolistTaskListPrefix, uid, page, pageSize)
	var resp []Task
	err := m.QueryCtx(ctx, &resp, key, func(conn *gorm.DB, v interface{}) error {
		err := conn.Debug().Model(&Task{}).Where("uid = ?", uid).Offset(int((page - 1) * pageSize)).Limit(int(pageSize)).Find(&resp).Error
		return err
	})
	return resp, err
}
