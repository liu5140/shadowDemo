package service

import (
	"github.com/shopspring/decimal"
)

type NoFoundAgentError struct {
	error
}

type LessThanDepositError struct {
	error
	Limit decimal.Decimal
}
