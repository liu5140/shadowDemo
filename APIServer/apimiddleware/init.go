package apimiddleware

import (
	"shadowDemo/zframework/logger"
	"strings"
)

var Log *logger.Logger

func init() {
	Log = logger.InitLog()
}

func IPResolve(ipString string) string {
	if ipString == "" {
		return ""
	}
	if strings.HasPrefix(ipString, "[::1]") {
		return "localhost"
	}
	return strings.Split(ipString, ":")[0]
}

const (
	ACTION_URL       = "/api/v1/action"
	BANK_LIST        = "query_banklist"
	DEPOSIT          = "deposit"
	WITHDRAW         = "withdraw"
	QUERY_DEPOSIT    = "query_order"
	QUERY_BALANCE    = "query_balance"
	VIEW_DEPOSIT_URL = "/view_deposit.shtml"
)
