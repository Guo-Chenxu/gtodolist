package logic

import (
	"context"
	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/pkg/errors"
	"gtodolist/app/task/cmd/rpc/internal/svc"
	"gtodolist/app/task/cmd/rpc/pb"
	"gtodolist/common/vo"
	"strconv"

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

	err = l.svcCtx.TaskModel.Delete(l.ctx, nil, tid)
	if err != nil && err != gormc.ErrNotFound {
		return nil, errors.Wrap(err, "删除任务失败")
	}

	return &pb.DeleteResp{
		Status:  vo.OK,
		Message: "删除成功",
	}, nil
}
