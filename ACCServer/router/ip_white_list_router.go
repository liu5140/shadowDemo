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

func iPWhiteListRouter(r *gin.RouterGroup) {
	// swagger:route POST /iPWhiteList iPWhiteList createIPWhiteList
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
	r.POST("/iPWhiteList", createIPWhiteList)

	// swagger:route POST /iPWhiteLists iPWhiteList searchIPWhiteList
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
	//       200: searchIPWhiteListResponse
	//       default: genericError
	r.POST("/iPWhiteLists", SearchIPWhiteList)

	// swagger:route GET /iPWhiteList iPWhiteList getIPWhiteList
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
	//       200: getIPWhiteListResponse
	//       default: genericError
	r.GET("/iPWhiteList", GetIPWhiteList)

	// swagger:route PUT /iPWhiteList iPWhiteList updateIPWhiteList
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
	r.PUT("/iPWhiteList", UpdateIPWhiteList)

	// swagger:route DELETE /iPWhiteList iPWhiteList deleteIPWhiteList
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
	r.DELETE("/iPWhiteList", DeleteIPWhiteList)

}

func createIPWhiteList(c *gin.Context) {
	iPWhiteListService := service.NewIPWhiteListService()
	request := &vo.CreateIPWhiteListRequest{}
	err := c.Bind(&request.Body)
	if err != nil {
		newClientError(c, err)
		return
	}
	profile := c.MustGet(PROFILE).(middleware.Profile)
	iPWhiteList := &do.IPWhiteList{
		CreatedBy:  profile.Account,
	}

	err = iPWhiteListService.CreateIPWhiteList(iPWhiteList)
	if err != nil {
		newServerError(c, err)
		return
	}

	newSuccess(c)
}

func SearchIPWhiteList(c *gin.Context) {
	iPWhiteListService := service.NewIPWhiteListService()
	request := &vo.SearchIPWhiteListRequest{}
	err := c.Bind(&request.Body)
	if err != nil {
		newClientError(c, err)
		return
	}

	condition := &dao.IPWhiteListSearchCondition{
		IDS:             request.Body.IDS,
		CreateStartTime: request.Body.CreateStartTime,
		CreateEndTime:   request.Body.CreateEndTime,
	}

	result, count, err := iPWhiteListService.SearchIPWhiteListPaging(condition, request.Body.PageNum, request.Body.PageSize)
	if err != nil {
		newServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, vo.SearchIPWhiteListResponseBody{
		Result: result,
		Count:  count,
	})
}

func GetIPWhiteList(c *gin.Context) {
	id, err := bindID(c)
	if err != nil {
		newClientError(c, err)
		return
	}
	iPWhiteListService := service.NewIPWhiteListService()
	iPWhiteList, err := iPWhiteListService.GetIPWhiteListByID(id)
	if err != nil {
		newServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, vo.GetIPWhiteListResponseBody{
		Result: *iPWhiteList,
	})
}

func UpdateIPWhiteList(c *gin.Context) {
	iPWhiteListService := service.NewIPWhiteListService()
	profile := c.MustGet(PROFILE).(middleware.Profile)
	request := &vo.UpdateIPWhiteListRequest{}
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

	err = iPWhiteListService.UpdateIPWhiteList(id, attrs)
	if err != nil {
		newServerError(c, err)
		return
	}
	newSuccess(c)
}

func DeleteIPWhiteList(c *gin.Context) {
	iPWhiteListService := service.NewIPWhiteListService()
	id, err := bindID(c)
	if err != nil {
		newClientError(c, err)
		return
	}
	err = iPWhiteListService.DeleteIPWhiteListByID(id)
	if err != nil {
		newServerError(c, err)
		return
	}
	newSuccess(c)
}
