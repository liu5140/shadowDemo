package vo

import (
	"errors"

	"github.com/shopspring/decimal"
)

type APIQueryBalanceReqData struct {
	RequestNo  string
	MerchantId string
}

type APIQueryBalanceResult struct {
	Balance       decimal.Decimal
	FrozenBalance decimal.Decimal
}

func (request APIQueryBalanceReqData) CheckRequestResolve() error {
	if request.MerchantId == "" {
		return errors.New("MerchantId not found")
	}
	return nil
}
