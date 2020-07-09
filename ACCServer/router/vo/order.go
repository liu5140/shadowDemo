package vo

import (
	"shadowDemo/model/do"
	"time"
)

// swagger:parameters order createOrder
type CreateOrderRequest struct {
	//in:body
	Body struct {
		
	}
}

// swagger:parameters order searchOrder
type SearchOrderRequest struct {
	//in:body
	Body struct {
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

//查询列表
//swagger:response searchOrderResponse
type SearchOrderResponse struct {
	//in: body
	Result SearchOrderResponseBody
}

// swagger:model
type SearchOrderResponseBody struct {
	Result []do.Order
	Count  int
}

//详情入参
// swagger:parameters order getOrder
type GetOrderRequest struct {
	//in: query
	ID int64
}

//详情出参
// swagger:response getOrderResponse
type GetOrderResponse struct {
	//in: body
	Result GetOrderResponseBody
}

//swagger:model
type GetOrderResponseBody struct {
	Result do.Order
}

//修改
// swagger:parameters order updateOrder
type UpdateOrderRequest struct {
	// in:query
	//id主键
	ID int64
	// in:body
	Body struct {
		
	}
}

//删除
// swagger:parameters order deleteOrder
type DeleteOrderRequest struct {
	//in: query
	ID int64
}
