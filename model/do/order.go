package do

import (
	"time"

	"github.com/shopspring/decimal"
)

type EnumPayType int

const (
	DefaultPay   EnumPayType = iota // 银联网关,(默认网银渠道,不需要卡号及卡主身份信息)
	FastUnionPay                    // 银联快捷 (需要卡号及卡主身份信息)
)

type Order struct {
	ID                int64 `gorm:"primary_key"`
	CreatedAt         *time.Time
	UpdatedAt         *time.Time
	DeletedAt         *time.Time
	CreatedBy         string
	OrderNo           string
	CustomerNo        string
	MerchantSecretKey string
	MerchantID        int64
	MerchantNo        string
	MerchantName      string
	MerchantOrderNo   string
	OrderPrice        decimal.Decimal `sql:"type:decimal(20,4);"` //订单金额
	OrderRealPrice    decimal.Decimal `sql:"type:decimal(20,4);"` //实际支付金额
	// OrderType           OperateType     // 充值=1，提现=2
	// OrderState          OrderState
	RelatedOrderNo string
	Remark         string
	Information    string
	Bank           string
	BankCardNo     string
	CardOwner      string
	//	PayType             EnumPayType
	CardOwnerCert  string // 卡主身份证号码
	CardOwnerPhone string // 卡主绑定手机号
	BankCardCvv    string // 信用卡cvv
	BankCardExpire string // 信用卡到期时间
	CreateTime     time.Time
	UpdateTime     time.Time
	ChannelCode    string
	RiskLevel      int
	PaymentID      int64
	//	Payment             BccPayment `gorm:"ForeignKey:PaymentID"`
	IsTest              bool
	MerchantCallbackURL string
	City                string
	Province            string
	Subbranch           string
	MerchantWithdrawFee decimal.Decimal `sql:"type:decimal(20,4);"`
	Operator            string          //订单分配的操作员
	OperatorNo          string          //订单分配的操作员帐号
	AccountNo           string          //订单分配的财务帐号
	AccountName         string          //订单分配的财务帐号名称
	RemitType           int             //出款类型，API代付/单笔代付
	Abnormal            bool            `gorm:"-"`
	ApiVersion          int
	CodeMerchantOrder   bool
	ClientIP            string //客户IP（view deposit IP）
}

func (order Order) TableName() string {
	return "order"
}
