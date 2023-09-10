package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"gtodolist/common/vo"

	"gtodolist/app/task/cmd/rpc/internal/svc"
	"gtodolist/app/task/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListTaskLogic {
	return &ListTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListTaskLogic) ListTask(in *pb.ListReq) (*pb.ListResp, error) {
	tasks, err := l.svcCtx.TaskModel.List(l.ctx, in.Uid, in.Start, in.Limit)
	if err != nil {
		return nil, errors.Wrap(vo.ErrDBerror, "数据库查询出错")
	}

	taskResp := make([]*pb.Task, len(tasks))
	for i := 0; i < len(tasks); i++ {
		task := tasks[i]
		taskResp[i] = &pb.Task{
			Id:        task.Id,
			Title:     task.Title,
			Content:   task.Content.String,
			Status:    int32(task.Status),
			CreateAt:  task.CreatedAt.Unix(),
			StartTime: task.StartTime.Unix(),
			EndTime:   task.EndTime.Time.Unix(),
		}
	}

	return &pb.ListResp{
		Status: vo.OK,
		Data: &pb.Data{
			Task:  taskResp,
			Total: int64(len(tasks)),
		},
		Message: vo.SUCCESS,
	}, nil
}

// DeleteListCache 删除 list 缓存
func (l *ListTaskLogic) DeleteListCache(uid int64) {
	prefix := "cache:gtodolist:task:list:"
	listKey := fmt.Sprintf("%s%v:*", prefix, uid)
	keys, _, _ := l.svcCtx.RedisClient.Scan(0, listKey, 0)

	for _, key := range keys {
		_, _ = l.svcCtx.RedisClient.Del(key)
	}
}
