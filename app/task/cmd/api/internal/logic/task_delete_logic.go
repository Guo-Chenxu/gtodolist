package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"gtodolist/app/task/cmd/rpc/pb"
	"gtodolist/common/ctxdata"
	"gtodolist/common/vo"

	"gtodolist/app/task/cmd/api/internal/svc"
	"gtodolist/app/task/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTaskDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskDeleteLogic {
	return &TaskDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TaskDeleteLogic) TaskDelete(req *types.DeleteReq) (resp *types.DeleteResp, err error) {
	deleteResp, err := l.svcCtx.TaskRpcClient.DeleteTask(l.ctx, &pb.DeleteReq{
		Uid: ctxdata.GetUidFromCtx(l.ctx),
		Id:  req.Id,
	})
	if err != nil {
		return &types.DeleteResp{
			Status:  int(vo.ErrServerCommonError.GetErrCode()),
			Message: err.Error(),
			Error:   err.Error(),
		}, nil
	}

	resp = &types.DeleteResp{}
	_ = copier.Copy(resp, deleteResp)
	return resp, err
}
