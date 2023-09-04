// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"fmt"
	"time"

	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var (
	cacheGtodolistTaskIdPrefix = "cache:gtodolist:task:id:"
)

type (
	taskModel interface {
		Insert(ctx context.Context, tx *gorm.DB, data *Task) error

		FindOne(ctx context.Context, id int64) (*Task, error)
		Update(ctx context.Context, tx *gorm.DB, data *Task) error

		Delete(ctx context.Context, tx *gorm.DB, id int64) error
		Transaction(ctx context.Context, fn func(db *gorm.DB) error) error
	}

	defaultTaskModel struct {
		gormc.CachedConn
		table string
	}

	Task struct {
		Id        int64          `gorm:"column:id"`               // 主键
		CreatedAt time.Time      `gorm:"column:created_at"`       // 创建时间
		UpdatedAt time.Time      `gorm:"column:updated_at"`       // 更新时间
		DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"` // 删除时间
		Uid       int64          `gorm:"column:uid"`              // 用户id
		Title     string         `gorm:"column:title"`            // 任务标题
		Status    int64          `gorm:"column:status"`           // 任务状态 0代办 1已完成
		Content   sql.NullString `gorm:"column:content"`          // 任务内容
		StartTime time.Time      `gorm:"column:start_time"`       // 任务开始时间
		EndTime   sql.NullTime   `gorm:"column:end_time"`         // 任务结束时间
	}
)

func (Task) TableName() string {
	return "`task`"
}

func newTaskModel(conn *gorm.DB, c cache.CacheConf) *defaultTaskModel {
	return &defaultTaskModel{
		CachedConn: gormc.NewConn(conn, c),
		table:      "`task`",
	}
}

func (m *defaultTaskModel) Insert(ctx context.Context, tx *gorm.DB, data *Task) error {

	err := m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Save(&data).Error
	}, m.getCacheKeys(data)...)
	return err
}

func (m *defaultTaskModel) FindOne(ctx context.Context, id int64) (*Task, error) {
	gtodolistTaskIdKey := fmt.Sprintf("%s%v", cacheGtodolistTaskIdPrefix, id)
	var resp Task
	err := m.QueryCtx(ctx, &resp, gtodolistTaskIdKey, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&Task{}).Where("`id` = ?", id).First(&resp).Error
	})
	switch err {
	case nil:
		return &resp, nil
	case gormc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTaskModel) Update(ctx context.Context, tx *gorm.DB, data *Task) error {
	old, err := m.FindOne(ctx, data.Id)
	if err != nil && err != ErrNotFound {
		return err
	}
	err = m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Save(data).Error
	}, m.getCacheKeys(old)...)
	return err
}

func (m *defaultTaskModel) getCacheKeys(data *Task) []string {
	if data == nil {
		return []string{}
	}
	gtodolistTaskIdKey := fmt.Sprintf("%s%v", cacheGtodolistTaskIdPrefix, data.Id)
	cacheKeys := []string{
		gtodolistTaskIdKey,
	}
	cacheKeys = append(cacheKeys, m.customCacheKeys(data)...)
	return cacheKeys
}

func (m *defaultTaskModel) Delete(ctx context.Context, tx *gorm.DB, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		if err == ErrNotFound {
			return nil
		}
		return err
	}
	err = m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Delete(&Task{}, id).Error
	}, m.getCacheKeys(data)...)
	return err
}

func (m *defaultTaskModel) Transaction(ctx context.Context, fn func(db *gorm.DB) error) error {
	return m.TransactCtx(ctx, fn)
}

func (m *defaultTaskModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheGtodolistTaskIdPrefix, primary)
}

func (m *defaultTaskModel) queryPrimary(conn *gorm.DB, v, primary interface{}) error {
	return conn.Model(&Task{}).Where("`id` = ?", primary).Take(v).Error
}
