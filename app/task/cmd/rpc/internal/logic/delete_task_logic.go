package logic

import (
	"context"
	"github.com/pkg/errors"
	"strconv"
	"time"

	"gtodolist/app/task/cmd/rpc/internal/svc"
	"gtodolist/app/task/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTaskLogic {
	return &DeleteTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteTaskLogic) DeleteTask(in *pb.DeleteReq) (*pb.DeleteResp, error) {
	tid, err := strconv.ParseInt(in.Id, 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "转换任务id失败, 请传入合法id")
	}
	DeleteListCache(l.svcCtx.RedisClient, in.Uid)

	go func() {
		l.Logger.Infof("删除任务缓存成功，uid为%d", in.Uid, "开始异步调用删除数据库")
		time.Sleep(time.Hour * 1)
		err := l.svcCtx.TaskModel.Delete(l.ctx, nil, tid)
		if err != nil {
			return
		}
	}()

	return &pb.DeleteResp{
		Status:  0,
		Message: "删除成功",
	}, nil
}
