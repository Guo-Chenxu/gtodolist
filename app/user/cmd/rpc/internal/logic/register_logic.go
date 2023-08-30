package logic

import (
	"context"
	"database/sql"
	"gtodolist/app/user/model"
	"gtodolist/common/tool"
	"gtodolist/common/vo"

	"gtodolist/app/user/cmd/rpc/internal/svc"
	"gtodolist/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterResp, error) {
	// 根据用户名查询用户是否存在
	user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, sql.NullString{
		String: in.Username,
		Valid:  true,
	})

	// 查询出错
	if err != nil && err != model.ErrNotFound {
		return &pb.RegisterResp{
			Status:  vo.ErrRequestParamError.GetErrCode(),
			Data:    "",
			Message: vo.ErrRequestParamError.GetErrMsg(),
			Error:   err.Error(),
		}, err
	}

	// 用户名已存在
	if user != nil {
		return &pb.RegisterResp{
			Status:  vo.ErrUserAlreadyRegisterError.GetErrCode(),
			Data:    "",
			Message: vo.ErrUserAlreadyRegisterError.GetErrMsg(),
			Error:   "",
		}, err
	}

	// 添加用户
	username := sql.NullString{
		String: in.Username,
		Valid:  true,
	}
	bp, _ := tool.BcryptByString(in.Password)
	passwordDigest := sql.NullString{
		String: bp,
		Valid:  true,
	}
	user = &model.User{
		Username:       username,
		PasswordDigest: passwordDigest,
	}
	err = l.svcCtx.UserModel.Insert(l.ctx, nil, user)

	// 注册失败
	if err != nil {
		return &pb.RegisterResp{
			Status:  vo.ErrDBerror.GetErrCode(),
			Data:    "",
			Message: vo.ErrDBerror.GetErrMsg(),
			Error:   err.Error(),
		}, err
	}

	return &pb.RegisterResp{
		Status:  vo.OK,
		Data:    vo.SUCCESS,
		Message: vo.SUCCESS,
		Error:   "",
	}, nil
}
