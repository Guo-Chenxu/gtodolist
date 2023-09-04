package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"gtodolist/app/task/cmd/rpc/pb"
	"gtodolist/common/vo"
	"strconv"

	"gtodolist/app/task/cmd/api/internal/svc"
	"gtodolist/app/task/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTaskCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskCreateLogic {
	return &TaskCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TaskCreateLogic) TaskCreate(req *types.CreateReq) (resp *types.CreateResp, err error) {
	status, err := strconv.Atoi(req.Status)
	if err != nil {
		status = 1
	}

	createResp, err := l.svcCtx.TaskRpcClient.CreateTask(l.ctx, &pb.CreateReq{
		Title:   req.Title,
		Content: req.Content,
		Status:  int32(status),
	})

	if err != nil {
		return &types.CreateResp{
			Status:  int(vo.ErrServerCommonError.GetErrCode()),
			Message: err.Error(),
			Error:   err.Error(),
		}, nil
	}

	resp = &types.CreateResp{}
	_ = copier.Copy(resp, createResp)
	return resp, err
}
