package vo

import "shadowDemo/zframework/logger"

var Log *logger.Logger

func init() {
	Log = logger.InitLog()
}

const (
	SUCCESS = "success"
)
const (
	PARAM_MERCHANT = "Merchant"
	PARAM_REQUEST  = "Request"
	PARAM_PAYTYPE  = "PayType"
	PARAM_RESPONSE = "Response"
	PARAM_ORDER    = "Order"
	PARAM_ENDPOINT = "Endpoint"
	PARAM_ERROR    = "Error"
)
