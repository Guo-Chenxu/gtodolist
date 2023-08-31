package logic

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"gtodolist/app/user/cmd/rpc/internal/svc"
	"gtodolist/app/user/cmd/rpc/pb"
	"gtodolist/app/user/cmd/rpc/userrpc"
	"gtodolist/app/user/model"
	"gtodolist/common/tool"
	"gtodolist/common/vo"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	// 查询用户是否存在且密码正确
	user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, sql.NullString{
		String: in.Username,
		Valid:  true,
	})
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrap(vo.ErrDBerror, "数据库查询出错")
	}
	if user == nil {
		return nil, errors.Wrap(vo.ErrUserNoExistsError, "用户不存在")
	}
	if !tool.CheckPasswordHash(in.Password, user.PasswordDigest.String) {
		return nil, errors.Wrap(vo.ErrUsernamePwdError, "密码匹配出错")
	}

	// 生成 token
	genToken := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := genToken.GenerateToken(&userrpc.GenerateTokenReq{
		UserId: user.Id,
	})
	if err != nil {
		return nil, errors.Wrap(vo.ErrGenerateTokenError, "生成 token 失败")
	}

	return &pb.LoginResp{
		Status: vo.OK,
		Data: &pb.Data{
			User: &pb.User{
				Id:       user.Id,
				Username: user.Username.String,
				CreateAt: user.CreatedAt.Unix(),
			},
			Token: tokenResp.AccessToken,
		},
		Message: vo.SUCCESS,
		Error:   "",
	}, nil
}
