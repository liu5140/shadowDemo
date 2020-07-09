package vo

import (
	"net/url"
	"shadowDemo/zframework/utils/encoder"

	"github.com/sirupsen/logrus"
)

//APIViewDepositRequest 请求
type APIViewDepositRequest struct {
	MerchantId   string `binding:"required"` //用户ID
	RequestNo    string `url:"-"`            //请求编号 (记录在log中)
	OrderNo      string `binding:"required"` //訂單编号
	EndpointCode string `url:"-"`            //三方
	Sign         string `binding:"required"` //签名
}

func (request APIViewDepositRequest) CalculateSign(secret string) APIViewDepositRequest {
	form := url.Values{}
	form.Add("MerchantId", request.MerchantId)
	form.Add("OrderNo", request.OrderNo)
	enStr := form.Encode() + "&Key=" + secret
	request.Sign = encoder.MD5(enStr)
	return request
}

func (request APIViewDepositRequest) VerifySign(secret string) bool {
	form := url.Values{}
	form.Add("MerchantId", request.MerchantId)
	form.Add("OrderNo", request.OrderNo)
	enStr := form.Encode() + "&Key=" + secret
	sign := encoder.MD5(enStr)

	Log.WithFields(logrus.Fields{
		"sign":         sign,
		"request_sign": request.Sign,
	}).Debug("APIViewDepositRequest.VerifySign")

	return sign == request.Sign
}

func (request APIViewDepositRequest) GetMerchantNo() string {
	return request.MerchantId
}
