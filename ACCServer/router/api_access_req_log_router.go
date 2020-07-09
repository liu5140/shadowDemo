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

func aPIAccessReqLogRouter(r *gin.RouterGroup) {
	// swagger:route POST /aPIAccessReqLog aPIAccessReqLog createAPIAccessReqLog
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
	r.POST("/aPIAccessReqLog", createAPIAccessReqLog)

	// swagger:route POST /aPIAccessReqLogs aPIAccessReqLog searchAPIAccessReqLog
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
	//       200: searchAPIAccessReqLogResponse
	//       default: genericError
	r.POST("/aPIAccessReqLogs", SearchAPIAccessReqLog)

	// swagger:route GET /aPIAccessReqLog aPIAccessReqLog getAPIAccessReqLog
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
	//       200: getAPIAccessReqLogResponse
	//       default: genericError
	r.GET("/aPIAccessReqLog", GetAPIAccessReqLog)

	// swagger:route PUT /aPIAccessReqLog aPIAccessReqLog updateAPIAccessReqLog
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
	r.PUT("/aPIAccessReqLog", UpdateAPIAccessReqLog)

	// swagger:route DELETE /aPIAccessReqLog aPIAccessReqLog deleteAPIAccessReqLog
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
	r.DELETE("/aPIAccessReqLog", DeleteAPIAccessReqLog)

}

func createAPIAccessReqLog(c *gin.Context) {
	aPIAccessReqLogService := service.NewAPIAccessReqLogService()
	request := &vo.CreateAPIAccessReqLogRequest{}
	err := c.Bind(&request.Body)
	if err != nil {
		newClientError(c, err)
		return
	}
	profile := c.MustGet(PROFILE).(middleware.Profile)
	aPIAccessReqLog := &do.APIAccessReqLog{
		CreatedBy:  profile.Account,
	}

	err = aPIAccessReqLogService.CreateAPIAccessReqLog(aPIAccessReqLog)
	if err != nil {
		newServerError(c, err)
		return
	}

	newSuccess(c)
}

func SearchAPIAccessReqLog(c *gin.Context) {
	aPIAccessReqLogService := service.NewAPIAccessReqLogService()
	request := &vo.SearchAPIAccessReqLogRequest{}
	err := c.Bind(&request.Body)
	if err != nil {
		newClientError(c, err)
		return
	}

	condition := &dao.APIAccessReqLogSearchCondition{
		IDS:             request.Body.IDS,
		CreateStartTime: request.Body.CreateStartTime,
		CreateEndTime:   request.Body.CreateEndTime,
	}

	result, count, err := aPIAccessReqLogService.SearchAPIAccessReqLogPaging(condition, request.Body.PageNum, request.Body.PageSize)
	if err != nil {
		newServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, vo.SearchAPIAccessReqLogResponseBody{
		Result: result,
		Count:  count,
	})
}

func GetAPIAccessReqLog(c *gin.Context) {
	id, err := bindID(c)
	if err != nil {
		newClientError(c, err)
		return
	}
	aPIAccessReqLogService := service.NewAPIAccessReqLogService()
	aPIAccessReqLog, err := aPIAccessReqLogService.GetAPIAccessReqLogByID(id)
	if err != nil {
		newServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, vo.GetAPIAccessReqLogResponseBody{
		Result: *aPIAccessReqLog,
	})
}

func UpdateAPIAccessReqLog(c *gin.Context) {
	aPIAccessReqLogService := service.NewAPIAccessReqLogService()
	profile := c.MustGet(PROFILE).(middleware.Profile)
	request := &vo.UpdateAPIAccessReqLogRequest{}
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

	err = aPIAccessReqLogService.UpdateAPIAccessReqLog(id, attrs)
	if err != nil {
		newServerError(c, err)
		return
	}
	newSuccess(c)
}

func DeleteAPIAccessReqLog(c *gin.Context) {
	aPIAccessReqLogService := service.NewAPIAccessReqLogService()
	id, err := bindID(c)
	if err != nil {
		newClientError(c, err)
		return
	}
	err = aPIAccessReqLogService.DeleteAPIAccessReqLogByID(id)
	if err != nil {
		newServerError(c, err)
		return
	}
	newSuccess(c)
}
