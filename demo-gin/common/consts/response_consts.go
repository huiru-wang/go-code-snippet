package consts

const (
	SuccessCode    int    = 200
	SuccessMsg     string = "Success"
	ParamErrorCode int    = 400
	ParmaErrorMsg  string = "Param error."
	FailCode       int    = 500
	FailMsg        string = "Server Error."

	NoAuthorizationCode  int    = 401
	NoAuthorizationMsg   string = "Token is empty, plz sign in."
	AuthorizationFailMsg string = "Validate token fail, plz sign in."
	CreateTokenErrorCode int    = 1000
	CreateTokenErrorMsg  string = "Create token failed, plz login again."

	RegisterErrorCode       int    = 10000
	RegisterErrorMsg        string = "Server error, register failed."
	DuplicateUserNameCode   int    = 10001
	DuplicateUserNameMsg    string = "User name is already exits."
	UserNameNotExitsCode    int    = 10003
	UserNameNotExitsMsg     string = "User name is not exits."
	UserPassNotMatchCode    int    = 10004
	UserPassNotMatchMsg     string = "User password is not match."
	UserPhoneUpdateFailCode int    = 10010
	UserPhoneUpdateFailMsg  string = "User phone update fail."
)
