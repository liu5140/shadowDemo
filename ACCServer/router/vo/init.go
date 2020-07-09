package vo

import (
	"shadowDemo/shadow-framework/logger"

	validator "gopkg.in/go-playground/validator.v9"
)

var Log *logger.Logger

var Validate *validator.Validate

const (
	TEMPLATE_PARAM_SUCCESS = "Success"
	TEMPLATE_PARAM_ERROR   = "Error"
)

func init() {
	Log = logger.InitLog()
}

// 错误消息
// swagger:response genericError
type GenericError struct {
	//in: body
	Body GenericMessageBody
}

// 成功消息
// swagger:response genericSuccess
type GenericSuccess struct {
	//in: body
	Body GenericMessageBody
}

// swagger:model
type GenericMessageBody struct {
	Msg  string
	Code int
}

// swagger:model
type EmptyMessageBody struct {
	Result []string
	Count  int
}
