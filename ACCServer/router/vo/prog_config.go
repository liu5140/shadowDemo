package vo

import (
	"shadowDemo/model/do"
	"time"
)

// swagger:parameters progConfig createProgConfig
type CreateProgConfigRequest struct {
	//in:body
	Body struct {
		//参数名称
		ParamName string
		//参数值
		ParamValue string
		//类型
		Type string
		//是否可用
		Disabled bool
		//是否需要加密 暂时没用到
		Encrypted bool
		//备注
		Comment string
	}
}

// swagger:parameters progConfig searchProgConfig
type SearchProgConfigRequest struct {
	//in:body
	Body struct {
		//参数名称
		ParamName string
		//类型
		Type string
		//id
		IDS []int64
		//创建开始时间
		CreateStartTime time.Time
		//创建截止时间
		CreateEndTime time.Time
		//页码
		PagingRequest
	}
}

//活动列表
//swagger:response searchProgConfigResponse
type SearchProgConfigResponse struct {
	//in: body
	Result SearchProgConfigResponseBody
}

// swagger:model
type SearchProgConfigResponseBody struct {
	Result []do.ProgConfig
	Count  int
}

// swagger:parameters progConfig getProgConfig
type GetProgConfigRequest struct {
	//in: query
	ID int64
}

// swagger:response getProgConfigResponse
type GetProgConfigResponse struct {
	//in: body
	Result GetProgConfigResponseBody
}

//swagger:model
type GetProgConfigResponseBody struct {
	Result do.ProgConfig
}

// swagger:parameters progConfig updateProgConfig
type UpdateProgConfigRequest struct {
	// in:query
	//id主键
	ID int64
	// in:body
	Body struct {
		//卡号
		CardNo string
		//开户银行
		BankName string
		//开户银行编码
		BankCode string
		//开户支行
		Branch string
		//开户城市
		City string
		//开户省份
		Province string
		//绑定邮箱
		Email string
		//绑定手机
		PhoneNumber string
		//开户时间
		OpeningDate *time.Time
		//备注
		Remark string
	}
}

// swagger:parameters progConfig deleteProgConfig
type DeleteProgConfigRequest struct {
	//in: query
	ID int64
}
