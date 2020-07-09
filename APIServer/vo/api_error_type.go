package vo

import "github.com/gin-gonic/gin"

const (
	RequestResolveFailed        gin.ErrorType = 401   //请求解析失败
	MerchantResolveFailed       gin.ErrorType = 402   //商户解析失败
	SginVerifyFailed            gin.ErrorType = 403   //请求校验失败
	IPVerifyFailed              gin.ErrorType = 404   //IP校验失败
	QueryChannelListFailed      gin.ErrorType = 9000  //查询通道失败
	QueryOrderFailed            gin.ErrorType = 9001  //查询订单失败
	WithdrawVerifyFailed        gin.ErrorType = 9002  //提现请求失败
	ThirdPartyFailed            gin.ErrorType = 10001 //三方支付请求失败
	ChannelIsNotSupported       gin.ErrorType = 10002 //请求的通道不被支持
	OrderVerifyFailed           gin.ErrorType = 10003 //订单验证失败
	MerchantInsufficientBalance gin.ErrorType = 10003 //商户余额不足
	OrderFailedApprove          gin.ErrorType = 10004 //提现订单被拒
	OrderCreateFailed           gin.ErrorType = 10005 //创建订单失败
	OrderRefuse                 gin.ErrorType = 10006 //拒绝支付
)
