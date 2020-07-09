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

func upmsRoleRouter(r *gin.RouterGroup) {
	// swagger:route POST /upmsRole upmsRole createUpmsRole
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
	r.POST("/upmsRole", createUpmsRole)

	// swagger:route POST /upmsRoles upmsRole searchUpmsRole
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
	//       200: searchUpmsRoleResponse
	//       default: genericError
	r.POST("/upmsRoles", SearchUpmsRole)

	// swagger:route GET /upmsRole upmsRole getUpmsRole
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
	//       200: getUpmsRoleResponse
	//       default: genericError
	r.GET("/upmsRole", GetUpmsRole)

	// swagger:route PUT /upmsRole upmsRole updateUpmsRole
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
	r.PUT("/upmsRole", UpdateUpmsRole)

	// swagger:route DELETE /upmsRole upmsRole deleteUpmsRole
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
	r.DELETE("/upmsRole", DeleteUpmsRole)

}

func createUpmsRole(c *gin.Context) {
	upmsRoleService := service.NewUpmsRoleService()
	request := &vo.CreateUpmsRoleRequest{}
	err := c.Bind(&request.Body)
	if err != nil {
		newClientError(c, err)
		return
	}
	profile := c.MustGet(PROFILE).(middleware.Profile)
	upmsRole := &do.UpmsRole{
		CreatedBy:  profile.Account,
	}

	err = upmsRoleService.CreateUpmsRole(upmsRole)
	if err != nil {
		newServerError(c, err)
		return
	}

	newSuccess(c)
}

func SearchUpmsRole(c *gin.Context) {
	upmsRoleService := service.NewUpmsRoleService()
	request := &vo.SearchUpmsRoleRequest{}
	err := c.Bind(&request.Body)
	if err != nil {
		newClientError(c, err)
		return
	}

	condition := &dao.UpmsRoleSearchCondition{
		IDS:             request.Body.IDS,
		CreateStartTime: request.Body.CreateStartTime,
		CreateEndTime:   request.Body.CreateEndTime,
	}

	result, count, err := upmsRoleService.SearchUpmsRolePaging(condition, request.Body.PageNum, request.Body.PageSize)
	if err != nil {
		newServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, vo.SearchUpmsRoleResponseBody{
		Result: result,
		Count:  count,
	})
}

func GetUpmsRole(c *gin.Context) {
	id, err := bindID(c)
	if err != nil {
		newClientError(c, err)
		return
	}
	upmsRoleService := service.NewUpmsRoleService()
	upmsRole, err := upmsRoleService.GetUpmsRoleByID(id)
	if err != nil {
		newServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, vo.GetUpmsRoleResponseBody{
		Result: *upmsRole,
	})
}

func UpdateUpmsRole(c *gin.Context) {
	upmsRoleService := service.NewUpmsRoleService()
	profile := c.MustGet(PROFILE).(middleware.Profile)
	request := &vo.UpdateUpmsRoleRequest{}
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

	err = upmsRoleService.UpdateUpmsRole(id, attrs)
	if err != nil {
		newServerError(c, err)
		return
	}
	newSuccess(c)
}

func DeleteUpmsRole(c *gin.Context) {
	upmsRoleService := service.NewUpmsRoleService()
	id, err := bindID(c)
	if err != nil {
		newClientError(c, err)
		return
	}
	err = upmsRoleService.DeleteUpmsRoleByID(id)
	if err != nil {
		newServerError(c, err)
		return
	}
	newSuccess(c)
}
