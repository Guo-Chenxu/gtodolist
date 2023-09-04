package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"gtodolist/app/task/cmd/rpc/pb"
	"gtodolist/common/ctxdata"
	"gtodolist/common/vo"
	"strconv"

	"gtodolist/app/task/cmd/api/internal/svc"
	"gtodolist/app/task/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTaskUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskUpdateLogic {
	return &TaskUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TaskUpdateLogic) TaskUpdate(req *types.UpdateReq) (resp *types.UpdateResp, err error) {
	status, err := strconv.Atoi(req.Status)
	if err != nil {
		status = 1
	}

	updateResp, err := l.svcCtx.TaskRpcClient.UpdateTask(l.ctx, &pb.UpdateReq{
		Id:      req.Id,
		Uid:     ctxdata.GetUidFromCtx(l.ctx),
		Title:   req.Title,
		Content: req.Title,
		Status:  int32(status),
	})
	if err != nil {
		return &types.UpdateResp{
			Status:  int(vo.ErrServerCommonError.GetErrCode()),
			Message: err.Error(),
			Error:   err.Error(),
		}, nil
	}

	resp = &types.UpdateResp{}
	_ = copier.Copy(resp, updateResp)
	return resp, err
}
