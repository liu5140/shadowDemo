package vo

import (
	"shadowDemo/model"
	"time"

	"github.com/shopspring/decimal"
)

// swagger:parameters session login
type LoginRequest struct {
	//in:body
	Body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
}

// 登录这类返回值
// swagger:response loginResponse
type LoginResponse struct {
	// in: body
	Result LoginResponseBody
}

//swagger:model
type LoginResponseBody struct {
	ID      int64
	Token   string
	Account string
	// 手机
	Mobile string

	NickName string
	// 生日
	Birthday *time.Time
	// 性别
	Sex model.Sex
	// 国家地区p
	Country string
	//在用户信息的接口里加个字段吧 hostIndex  (不知道干啥)
	HostIndex int
	//代号
	Code           int
	Balance        decimal.Decimal
	IsSecurePwdSet bool
	Recommend      string
}

// swagger:parameters  session createSecureToken
type CreateSecureTokenRequest struct {
	//in:body
	Body struct {
		Password string
	}
}

// 创建安全令牌
// swagger:response CreateSecureTokenResponse
type CreateSecureTokenResponse struct {
	// in: body
	Result CreateSecureTokenResponseBody
}

// swagger:model
type CreateSecureTokenResponseBody struct {
	Token string
}
