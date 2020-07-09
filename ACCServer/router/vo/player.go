package vo

import (
	"shadowDemo/model"
	"time"

	"github.com/shopspring/decimal"
)

// swagger:parameters player createPlayer
type CreatePlayerRequest struct {
	//in:body
	Body struct {
		// 帐号
		PhoneCode string
		// 帐号
		Account string `binding:"required"`
		// 登录密码
		LoginPassword string
		// 推荐人的推广码, (如果玩家是被推荐的）
		Recommend string
	}
}

// swagger:parameters payment createDeposit
type CreateDepositRequest struct {
	//钱
	Money string `binding:"required"`
	//通道
	CoinName string `binding:"required"`
	//交易密码
	TransactionPassword string ` binding:"required"`
	//单笔充值还是会员卡充值(1=单笔 2=会员卡)
	DepositType string `json:"depositType"`
}

// swagger:parameters playerInformation changeSecurePassword
type ChangeDsSecurePasswordRequest struct {
	NewPassword string `binding:"required" `
	OldPassword string `binding:"required" `
}

// swagger:parameters playerInformation setNickName
type SetNickNameRequest struct {
	// in: body
	NickName string `binding:"required"`
}

type SetAccountRequest struct {
	// in: body
	Account string `binding:"required"`
}

type SmsAccountRequest struct {
	// in: body
	UserName string `binding:"required"`
}

//swagger:model
type PlayerResponseBody struct {
	ID      int64
	Account string
	// 手机
	Mobile   string
	NickName string
	// 生日
	Birthday *time.Time
	// 性别
	Sex model.Sex
	// 国家地区p
	Country string
	//代号
	Code             int //余额
	Balance          decimal.Decimal
	IsSecurePwdSet   bool
	Recommend        string
	IncomeTotal      decimal.Decimal
	TodayIncomeTotal string
}

type SelectDepositOrderRequest struct {
	// in: body
	OrderNo string
	PagingRequest
}

type CoinNameRequest struct {
	CoinName string `binding:"required"`
}

type FundChangesRequest struct {
	// in: body
	PagingRequest
}

type SzFundChangesRequest struct {
	// in: body
	PagingRequest
}

type WebSocketRequest struct {
	Message string `binding:"required" `
}
