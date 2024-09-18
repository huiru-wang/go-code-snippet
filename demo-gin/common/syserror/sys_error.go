package syserror

import "go-gin/common/consts"

var RegisterError = &SysError{ErrorCode: consts.RegisterErrorCode, ErrorMsg: consts.RegisterErrorMsg}
var SystemError = &SysError{ErrorCode: consts.FailCode, ErrorMsg: consts.FailMsg}
var ParamError = &SysError{ErrorCode: consts.ParamErrorCode, ErrorMsg: consts.ParmaErrorMsg}
var UserNameExits = &SysError{ErrorCode: consts.DuplicateUserNameCode, ErrorMsg: consts.DuplicateUserNameMsg}
var UserNameNotExits = &SysError{ErrorCode: consts.UserNameNotExitsCode, ErrorMsg: consts.UserNameNotExitsMsg}
var UserPassNotMatch = &SysError{ErrorCode: consts.UserPassNotMatchCode, ErrorMsg: consts.UserPassNotMatchMsg}
var AuthError = &SysError{ErrorCode: consts.NoAuthorizationCode, ErrorMsg: consts.NoAuthorizationMsg}
var CreateTokenError = &SysError{ErrorCode: consts.CreateTokenErrorCode, ErrorMsg: consts.CreateTokenErrorMsg}

type SysError struct {
	ErrorCode int
	ErrorMsg  string
}

func Error(errorCode int, errorMsg string) *SysError {
	return &SysError{ErrorCode: errorCode, ErrorMsg: errorMsg}
}
