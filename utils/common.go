package utils

// RetCodeModel RetCodeModel
type RetCodeModel struct {
	Success                 int
	ExceptionError          int
	HTTPBadRequest          int
	HTTPUnAuthorized        int
	HTTPNotFound            int
	HTTPInternalServerError int
	ParamRequired           int
	NotFoundInfo            int
	ParamError              int
	LoginError              int
	VerifyFailed            int
}

var (

	// RetCode RetCode
	RetCode = RetCodeModel{
		Success:                 200,
		HTTPBadRequest:          400,
		HTTPUnAuthorized:        401,
		HTTPNotFound:            404,
		HTTPInternalServerError: 500,
		ParamRequired:           1001,
		NotFoundInfo:            1002,
		ParamError:              1003,
		LoginError:              1004,
		VerifyFailed:            1005,
	}
	// ErrorCodeMessage ErrorCodeMessage
	ErrorCodeMessage = map[int]string{
		200:  "SUCCESS",
		401:  "登录失败",
		500:  "异常错误",
		1001: "缺少参数",
		1002: "参数不全",
		1003: "参数错误",
		1004: "未登录",
		1005: "权限验证失败",
	}
)
