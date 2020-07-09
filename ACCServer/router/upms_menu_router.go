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

func upmsMenuRouter(r *gin.RouterGroup) {
	// swagger:route POST /upmsMenu upmsMenu createUpmsMenu
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
	r.POST("/upmsMenu", createUpmsMenu)

	// swagger:route POST /upmsMenus upmsMenu searchUpmsMenu
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
	//       200: searchUpmsMenuResponse
	//       default: genericError
	r.POST("/upmsMenus", SearchUpmsMenu)

	// swagger:route GET /upmsMenu upmsMenu getUpmsMenu
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
	//       200: getUpmsMenuResponse
	//       default: genericError
	r.GET("/upmsMenu", GetUpmsMenu)

	// swagger:route PUT /upmsMenu upmsMenu updateUpmsMenu
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
	r.PUT("/upmsMenu", UpdateUpmsMenu)

	// swagger:route DELETE /upmsMenu upmsMenu deleteUpmsMenu
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
	r.DELETE("/upmsMenu", DeleteUpmsMenu)

}

func createUpmsMenu(c *gin.Context) {
	upmsMenuService := service.NewUpmsMenuService()
	request := &vo.CreateUpmsMenuRequest{}
	err := c.Bind(&request.Body)
	if err != nil {
		newClientError(c, err)
		return
	}
	profile := c.MustGet(PROFILE).(middleware.Profile)
	upmsMenu := &do.UpmsMenu{
		CreatedBy:  profile.Account,
	}

	err = upmsMenuService.CreateUpmsMenu(upmsMenu)
	if err != nil {
		newServerError(c, err)
		return
	}

	newSuccess(c)
}

func SearchUpmsMenu(c *gin.Context) {
	upmsMenuService := service.NewUpmsMenuService()
	request := &vo.SearchUpmsMenuRequest{}
	err := c.Bind(&request.Body)
	if err != nil {
		newClientError(c, err)
		return
	}

	condition := &dao.UpmsMenuSearchCondition{
		IDS:             request.Body.IDS,
		CreateStartTime: request.Body.CreateStartTime,
		CreateEndTime:   request.Body.CreateEndTime,
	}

	result, count, err := upmsMenuService.SearchUpmsMenuPaging(condition, request.Body.PageNum, request.Body.PageSize)
	if err != nil {
		newServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, vo.SearchUpmsMenuResponseBody{
		Result: result,
		Count:  count,
	})
}

func GetUpmsMenu(c *gin.Context) {
	id, err := bindID(c)
	if err != nil {
		newClientError(c, err)
		return
	}
	upmsMenuService := service.NewUpmsMenuService()
	upmsMenu, err := upmsMenuService.GetUpmsMenuByID(id)
	if err != nil {
		newServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, vo.GetUpmsMenuResponseBody{
		Result: *upmsMenu,
	})
}

func UpdateUpmsMenu(c *gin.Context) {
	upmsMenuService := service.NewUpmsMenuService()
	profile := c.MustGet(PROFILE).(middleware.Profile)
	request := &vo.UpdateUpmsMenuRequest{}
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

	err = upmsMenuService.UpdateUpmsMenu(id, attrs)
	if err != nil {
		newServerError(c, err)
		return
	}
	newSuccess(c)
}

func DeleteUpmsMenu(c *gin.Context) {
	upmsMenuService := service.NewUpmsMenuService()
	id, err := bindID(c)
	if err != nil {
		newClientError(c, err)
		return
	}
	err = upmsMenuService.DeleteUpmsMenuByID(id)
	if err != nil {
		newServerError(c, err)
		return
	}
	newSuccess(c)
}
