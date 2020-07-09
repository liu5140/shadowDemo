package model

type (
	Device       int
	AccountState int
)

const (
	_         Device = iota
	PC               //来源终端: 电脑       1
	Mobile           //来源终端：手机端 	2
	H5               //来源终端：手机端H5 	3
	IOS              //来源终端: 苹果手机 	4
	Android          //来源终端: 安卓手机 	5
	AllDevice        //全部	6
)

const (
	_ AccountState = iota
	// 正常		1
	Normal
	// 帐号冻结	2
	Frozen
	// 余额冻结	3
	BalanceFrozen
	// 帐号停用	4
	Deleted
	// 审核中		5
	Auditing
	// 审核不通过		6
	AuditRefuse
)