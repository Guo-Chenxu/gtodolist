package vo

const (
	//状态码
	OK           int32 = 200
	NO_REQUEST   int32 = 400
	WRONG_SERVER int32 = 500
	//状态信息
	SUCCESS string = "成功执行！！！！"

	//服务错误码
	SERVER_COMMON_ERROR int32 = 100001
	//请求参数错误
	REUQEST_PARAM_ERROR int32 = 100002
	//token过期
	TOKEN_EXPIRE_ERROR int32 = 100003
	//token生成错误
	TOKEN_GENERATE_ERROR int32 = 100004
	//数据库错误
	DB_ERROR int32 = 100005
	//数据库更新错误
	DB_UPDATE_AFFECTED_ZERO_ERROR int32 = 100006
)

var (
	ErrGenerateTokenError       = NewErrMsg("生成token失败", NO_REQUEST)
	ErrUsernamePwdError         = NewErrMsg("账号或密码不正确", NO_REQUEST)
	ErrUserNoExistsError        = NewErrMsg("用户不存在", NO_REQUEST)
	ErrUserAlreadyRegisterError = NewErrMsg("该用户名已被注册", NO_REQUEST)
	ErrDBerror                  = NewErrMsg("数据库错误")
	ErrServerCommonError        = NewErrMsg("服务器请求错误")
	ErrRequestParamError        = NewErrMsg("请求参数错误", NO_REQUEST)
	ErrTokenExpireError         = NewErrMsg("Token已经过期")
	ErrUpdateAffectedZeroError  = NewErrMsg("数据库更新错误")
)
