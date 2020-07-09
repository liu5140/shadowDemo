package vo

import "errors"

type APIChannelReqData struct {
	UserLevel string `binding:"required"`
}

type APIChannelResult struct {
	Data []APIChannelResData
}

type APIChannelResData struct {
	BankCode string //银行Code
	Name     string
	Max      string // 单位元
	Min      string
}

func (request APIChannelReqData) CheckRequestResolve() error {
	if request.UserLevel == "" {
		return errors.New("UserLevel is not found")
	}
	return nil
}
