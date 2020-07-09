package vo

//APIDepositCallbackRequest 充值回调请求
type APIDepositCallbackRequest struct {
	MerchantId string
	RequestId  string
	OrderNo    string
	Money      string
	Code       int64
	Time       string
	Message    string
	Type       int64
	Fee        string
	Sign       string `url:"-"`
}
