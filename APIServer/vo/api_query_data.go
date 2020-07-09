package vo

import "errors"

type APIQueryDepositReqData struct {
	RequestNo      string
	RequestOrderNo string //用户订单号
	OrderNo        string //平台订单号
}

type APIQueryDepositResult struct {
	OrderNo        string
	RequestOrderNo string
	OrderStatus    string
	Money          string
	CreateTime     string
}

func (request APIQueryDepositReqData) CheckRequestResolve() error {
	if request.RequestOrderNo == "" && request.OrderNo == "" {
		return errors.New("RequestOrderNo and OrderNo are not found")
	}
	return nil
}
