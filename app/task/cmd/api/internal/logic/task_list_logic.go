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

type TaskListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTaskListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskListLogic {
	return &TaskListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TaskListLogic) TaskList(req *types.ListReq) (resp *types.ListResp, err error) {
	// 准备参数, 规定默认 limit 为 20, start 为 1
	limit, err := strconv.Atoi(req.Limit)
	if err != nil || limit <= 0 || limit > 20 {
		limit = 20
	}
	start, err := strconv.Atoi(req.Start)
	if err != nil || start <= 0 {
		start = 1
	}
	uid := ctxdata.GetUidFromCtx(l.ctx)

	listResp, err := l.svcCtx.TaskRpcClient.ListTask(l.ctx, &pb.ListReq{
		Uid:   uid,
		Limit: int64(limit),
		Start: int64(start),
	})

	if err != nil {
		return &types.ListResp{
			Status:  int(vo.ErrServerCommonError.GetErrCode()),
			Message: err.Error(),
			Error:   err.Error(),
		}, nil
	}

	resp = &types.ListResp{}
	_ = copier.Copy(resp, listResp)
	return resp, err
}
