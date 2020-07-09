package vo

import (
	"errors"
)

type APIDepositReqData struct {
	RequestNo      string
	RequestOrderNo string `binding:"required"`
	BankCode       string `binding:"required"` //银行Code
	Money          string `binding:"required"` //订单金额
	CallBackUrl    string //回调地址
	UserLevel      string `binding:"required"` //用户风险级别
	Message        string // 附加信息
	PayType        string // 默认0, 银联网关, 1=银联快捷
	CardNo         string // 银行卡号
	OwnerName      string // 银行账户姓名
	OwnerPhone     string // 绑定手机号码
	OwnerId        string // 绑定身份证号
	CardCvv        string // 信用卡后三位
	CardExpireDate string // 信用卡有效期
}

type APIDepositResult struct {
	OrderNo        string
	RequestOrderNo string
	Url            string
	Message        string
}

func (request APIDepositReqData) CheckRequestResolve() error {
	if request.RequestOrderNo == "" {
		return errors.New("RequestOrderNo is not found")
	}
	if request.BankCode == "" {
		return errors.New("BankCode is not found")
	}
	if request.Money == "" {
		return errors.New("Money is not found")
	}
	if request.UserLevel == "" {
		return errors.New("UserLevel is not found")
	}
	return nil
}
