package logic

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"gtodolist/app/task/model"
	"gtodolist/common/vo"
	"strconv"
	"time"

	"gtodolist/app/task/cmd/rpc/internal/svc"
	"gtodolist/app/task/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTaskLogic {
	return &UpdateTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateTaskLogic) UpdateTask(in *pb.UpdateReq) (*pb.UpdateResp, error) {
	// 先根据 id 查询该用户是否存在该任务
	id, err := strconv.Atoi(in.Id)
	if err != nil {
		return nil, errors.Wrap(vo.ErrRequestParamError, "任务id错误")
	}

	task, err := l.svcCtx.TaskModel.FindOne(l.ctx, int64(id))
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrap(vo.ErrDBerror, "数据库查询出错")
	}

	if task.Uid != in.Uid {
		return nil, errors.Wrap(vo.ErrRequestParamError, "没有这条任务信息")
	}

	// 删除list缓存
	NewListTaskLogic(l.ctx, l.svcCtx).DeleteListCache(in.Uid)

	// 修改任务
	task = &model.Task{
		Id:        task.Id,
		CreatedAt: task.CreatedAt,
		Uid:       task.Uid,
		Title:     in.Title,
		Status:    int64(in.Status),
		Content: sql.NullString{
			String: in.Content,
			Valid:  true,
		},
		StartTime: time.Now(),
	}
	err = l.svcCtx.TaskModel.Update(l.ctx, nil, task)
	if err != nil {
		return nil, errors.Wrap(vo.ErrDBerror, "数据库修改失败")
	}

	return &pb.UpdateResp{
		Status:  vo.OK,
		Data:    vo.SUCCESS,
		Message: vo.SUCCESS,
	}, nil
}
