package vo

import "errors"

type APIWithdrawReqData struct {
	RequestNo      string
	RequestOrderNo string `binding:"required"` //订单ID
	CardNo         string `binding:"required"` //用户卡号
	CardName       string `binding:"required"` //银行卡开卡人姓名
	BankCode       string `binding:"required"` //银行Code
	CallBackUrl    string //回调地址
	Money          string `binding:"required"` //订单金额
	UserLevel      string //用户风险级别
	Message        string //附加信息
	Province       string `binding:"required"` //银行卡开卡省份
	City           string `binding:"required"` //银行卡开卡城市
	Branch         string `binding:"required"` //支行信息
	BankPhone      string
	BankBranchCode string
}

type APIWithdrawResult struct {
	OrderNo        string
	RequestOrderNo string
	Message        string
}

func (request APIWithdrawReqData) CheckRequestResolve() error {
	if request.RequestOrderNo == "" {
		return errors.New("RequestOrderNo is not found")
	}
	if request.CardNo == "" {
		return errors.New("CardNo is not found")
	}
	if request.CardName == "" {
		return errors.New("CardName is not found")
	}
	if request.BankCode == "" {
		return errors.New("BankCode is not found")
	}
	if request.Money == "" {
		return errors.New("Money is not found")
	}
	if request.Province == "" {
		return errors.New("Province is not found")
	}
	if request.City == "" {
		return errors.New("City is not found")
	}
	if request.Branch == "" {
		return errors.New("Branch is not found")
	}
	return nil
}
