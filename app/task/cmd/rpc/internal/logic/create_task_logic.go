package logic

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"gtodolist/app/task/model"
	"gtodolist/common/vo"
	"time"

	"gtodolist/app/task/cmd/rpc/internal/svc"
	"gtodolist/app/task/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTaskLogic {
	return &CreateTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateTaskLogic) CreateTask(in *pb.CreateReq) (*pb.CreateResp, error) {
	task := &model.Task{
		Uid:    in.Uid,
		Title:  in.Title,
		Status: int64(in.Status),
		Content: sql.NullString{
			String: in.Content,
			Valid:  true,
		},
		StartTime: time.Now(),
	}

	err := l.svcCtx.TaskModel.Insert(l.ctx, nil, task)
	if err != nil {
		return nil, errors.Wrap(vo.ErrDBerror, "数据库插入出错")
	}

	return &pb.CreateResp{
		Status:  vo.OK,
		Data:    vo.SUCCESS,
		Message: vo.SUCCESS,
	}, nil
}
