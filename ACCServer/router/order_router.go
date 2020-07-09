package router

import (
	"net/http"
	"shadowDemo/ACCServer/router/vo"
	"shadowDemo/middleware"
	"shadowDemo/model/dao"
	"shadowDemo/model/do"
	"shadowDemo/service"

	"github.com/gin-gonic/gin"
)

func orderRouter(r *gin.RouterGroup) {
	// swagger:route POST /order order createOrder
	//
	// 创建;
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http
	//
	//     Security:
	//       api_key:
	//       oauth: read, write
	//
	//     Responses:
	//       200: genericSuccess
	//       default: genericError
	r.POST("/order", createOrder)

	// swagger:route POST /orders order searchOrder
	//
	// 查询列表;
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http
	//
	//     Security:
	//       api_key:
	//       oauth: read, write
	//
	//     Responses:
	//       200: searchOrderResponse
	//       default: genericError
	r.POST("/orders", SearchOrder)

	// swagger:route GET /order order getOrder
	//
	// 详情;
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http
	//
	//     Security:
	//       api_key:
	//       oauth: read, write
	//
	//     Responses:
	//       200: getOrderResponse
	//       default: genericError
	r.GET("/order", GetOrder)

	// swagger:route PUT /order order updateOrder
	//
	// 修改;
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http
	//
	//     Security:
	//       api_key:
	//       oauth: read, write
	//
	//     Responses:
	//       200: genericSuccess
	//       default: genericError
	r.PUT("/order", UpdateOrder)

	// swagger:route DELETE /order order deleteOrder
	//
	// 删除;
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http
	//
	//     Security:
	//       api_key:
	//       oauth: read, write
	//
	//     Responses:
	//       200: genericSuccess
	//       default: genericError
	r.DELETE("/order", DeleteOrder)

}

func createOrder(c *gin.Context) {
	orderService := service.NewOrderService()
	request := &vo.CreateOrderRequest{}
	err := c.Bind(&request.Body)
	if err != nil {
		newClientError(c, err)
		return
	}
	profile := c.MustGet(PROFILE).(middleware.Profile)
	order := &do.Order{
		CreatedBy:  profile.Account,
	}

	err = orderService.CreateOrder(order)
	if err != nil {
		newServerError(c, err)
		return
	}

	newSuccess(c)
}

func SearchOrder(c *gin.Context) {
	orderService := service.NewOrderService()
	request := &vo.SearchOrderRequest{}
	err := c.Bind(&request.Body)
	if err != nil {
		newClientError(c, err)
		return
	}

	condition := &dao.OrderSearchCondition{
		IDS:             request.Body.IDS,
		CreateStartTime: request.Body.CreateStartTime,
		CreateEndTime:   request.Body.CreateEndTime,
	}

	result, count, err := orderService.SearchOrderPaging(condition, request.Body.PageNum, request.Body.PageSize)
	if err != nil {
		newServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, vo.SearchOrderResponseBody{
		Result: result,
		Count:  count,
	})
}

func GetOrder(c *gin.Context) {
	id, err := bindID(c)
	if err != nil {
		newClientError(c, err)
		return
	}
	orderService := service.NewOrderService()
	order, err := orderService.GetOrderByID(id)
	if err != nil {
		newServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, vo.GetOrderResponseBody{
		Result: *order,
	})
}

func UpdateOrder(c *gin.Context) {
	orderService := service.NewOrderService()
	profile := c.MustGet(PROFILE).(middleware.Profile)
	request := &vo.UpdateOrderRequest{}
	id, err := bindID(c)
	if err != nil {
		newClientError(c, err)
		return
	}
	if err := c.Bind(&request.Body); err != nil {
		newClientError(c, err)
		return
	}
	attrs := map[string]interface{}{}
	//修改人
	attrs["created_by"] = profile.Account

	err = orderService.UpdateOrder(id, attrs)
	if err != nil {
		newServerError(c, err)
		return
	}
	newSuccess(c)
}

func DeleteOrder(c *gin.Context) {
	orderService := service.NewOrderService()
	id, err := bindID(c)
	if err != nil {
		newClientError(c, err)
		return
	}
	err = orderService.DeleteOrderByID(id)
	if err != nil {
		newServerError(c, err)
		return
	}
	newSuccess(c)
}
