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

func aPIAccessResLogRouter(r *gin.RouterGroup) {
	// swagger:route POST /aPIAccessResLog aPIAccessResLog createAPIAccessResLog
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
	r.POST("/aPIAccessResLog", createAPIAccessResLog)

	// swagger:route POST /aPIAccessResLogs aPIAccessResLog searchAPIAccessResLog
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
	//       200: searchAPIAccessResLogResponse
	//       default: genericError
	r.POST("/aPIAccessResLogs", SearchAPIAccessResLog)

	// swagger:route GET /aPIAccessResLog aPIAccessResLog getAPIAccessResLog
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
	//       200: getAPIAccessResLogResponse
	//       default: genericError
	r.GET("/aPIAccessResLog", GetAPIAccessResLog)

	// swagger:route PUT /aPIAccessResLog aPIAccessResLog updateAPIAccessResLog
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
	r.PUT("/aPIAccessResLog", UpdateAPIAccessResLog)

	// swagger:route DELETE /aPIAccessResLog aPIAccessResLog deleteAPIAccessResLog
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
	r.DELETE("/aPIAccessResLog", DeleteAPIAccessResLog)

}

func createAPIAccessResLog(c *gin.Context) {
	aPIAccessResLogService := service.NewAPIAccessResLogService()
	request := &vo.CreateAPIAccessResLogRequest{}
	err := c.Bind(&request.Body)
	if err != nil {
		newClientError(c, err)
		return
	}
	profile := c.MustGet(PROFILE).(middleware.Profile)
	aPIAccessResLog := &do.APIAccessResLog{
		CreatedBy:  profile.Account,
	}

	err = aPIAccessResLogService.CreateAPIAccessResLog(aPIAccessResLog)
	if err != nil {
		newServerError(c, err)
		return
	}

	newSuccess(c)
}

func SearchAPIAccessResLog(c *gin.Context) {
	aPIAccessResLogService := service.NewAPIAccessResLogService()
	request := &vo.SearchAPIAccessResLogRequest{}
	err := c.Bind(&request.Body)
	if err != nil {
		newClientError(c, err)
		return
	}

	condition := &dao.APIAccessResLogSearchCondition{
		IDS:             request.Body.IDS,
		CreateStartTime: request.Body.CreateStartTime,
		CreateEndTime:   request.Body.CreateEndTime,
	}

	result, count, err := aPIAccessResLogService.SearchAPIAccessResLogPaging(condition, request.Body.PageNum, request.Body.PageSize)
	if err != nil {
		newServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, vo.SearchAPIAccessResLogResponseBody{
		Result: result,
		Count:  count,
	})
}

func GetAPIAccessResLog(c *gin.Context) {
	id, err := bindID(c)
	if err != nil {
		newClientError(c, err)
		return
	}
	aPIAccessResLogService := service.NewAPIAccessResLogService()
	aPIAccessResLog, err := aPIAccessResLogService.GetAPIAccessResLogByID(id)
	if err != nil {
		newServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, vo.GetAPIAccessResLogResponseBody{
		Result: *aPIAccessResLog,
	})
}

func UpdateAPIAccessResLog(c *gin.Context) {
	aPIAccessResLogService := service.NewAPIAccessResLogService()
	profile := c.MustGet(PROFILE).(middleware.Profile)
	request := &vo.UpdateAPIAccessResLogRequest{}
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

	err = aPIAccessResLogService.UpdateAPIAccessResLog(id, attrs)
	if err != nil {
		newServerError(c, err)
		return
	}
	newSuccess(c)
}

func DeleteAPIAccessResLog(c *gin.Context) {
	aPIAccessResLogService := service.NewAPIAccessResLogService()
	id, err := bindID(c)
	if err != nil {
		newClientError(c, err)
		return
	}
	err = aPIAccessResLogService.DeleteAPIAccessResLogByID(id)
	if err != nil {
		newServerError(c, err)
		return
	}
	newSuccess(c)
}
