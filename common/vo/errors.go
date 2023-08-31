package vo

import "fmt"

/**
常用通用固定错误
*/

type CodeError struct {
	errCode int32
	errMsg  string
}

// 返回给前端的错误码
func (e *CodeError) GetErrCode() int32 {
	return e.errCode
}

// 返回给前端显示端错误信息
func (e *CodeError) GetErrMsg() string {
	return e.errMsg
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s", e.errCode, e.errMsg)
}

func NewErrMsg(errMsg string, errCode ...int32) *CodeError {
	if len(errCode) > 0 {
		return &CodeError{errCode: errCode[0], errMsg: errMsg}
	}
	return &CodeError{errCode: SERVER_COMMON_ERROR, errMsg: errMsg}
}
