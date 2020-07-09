package router

import (
	"net/http"
	"shadowDemo/ACCServer/router/vo"
	"shadowDemo/middleware"
	"shadowDemo/model/dao"
	"shadowDemo/model/do"
	"shadowDemo/service"
	"strings"

	"github.com/gin-gonic/gin"
)

func progConfigRouter(r *gin.RouterGroup) {
	// swagger:route POST /progConfig progConfig createProgConfig
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
	r.POST("/progConfig", createProgConfig)

	// swagger:route POST /progConfig progConfig searchProgConfig
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
	//       200: searchProgConfigResponse
	//       default: genericError
	r.POST("/progConfigs", SearchProgConfig)

	// swagger:route GET /progConfig progConfig getProgConfig
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
	//       200: getProgConfigResponse
	//       default: genericError
	r.GET("/progConfig", GetProgConfig)

	// swagger:route PUT /progConfig progConfig updateProgConfig
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
	r.PUT("/progConfig", UpdateProgConfig)

	// swagger:route PUT /progConfig progConfig deleteProgConfig
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
	r.DELETE("/progConfig", DeleteProgConfig)

}

func createProgConfig(c *gin.Context) {
	progConfigService := service.NewProgConfigService()
	request := &vo.CreateProgConfigRequest{}
	err := c.Bind(&request.Body)
	if err != nil {
		newClientError(c, err)
		return
	}
	profile := c.MustGet(PROFILE).(middleware.Profile)
	progConfig := &do.ProgConfig{
		CreatedBy:  profile.Account,
		ParamName:  strings.TrimSpace(request.Body.ParamName),
		ParamValue: strings.TrimSpace(request.Body.ParamValue),
		Type:       strings.TrimSpace(request.Body.Type),
		Disabled:   request.Body.Disabled,
		Encrypted:  request.Body.Encrypted,
		Comment:    strings.TrimSpace(request.Body.Comment),
	}

	err = progConfigService.CreateProgConfig(progConfig)
	if err != nil {
		newServerError(c, err)
		return
	}

	newSuccess(c)
}

func SearchProgConfig(c *gin.Context) {
	progConfigService := service.NewProgConfigService()
	request := &vo.SearchProgConfigRequest{}
	err := c.Bind(&request.Body)
	if err != nil {
		newClientError(c, err)
		return
	}

	condition := &dao.ProgConfigSearchCondition{
		ParamName:       request.Body.ParamName,
		Type:            request.Body.Type,
		IDS:             request.Body.IDS,
		CreateStartTime: request.Body.CreateStartTime,
		CreateEndTime:   request.Body.CreateEndTime,
	}

	result, count, err := progConfigService.SearchProgConfigPaging(condition, request.Body.PageNum, request.Body.PageSize)
	if err != nil {
		newServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, vo.SearchProgConfigResponseBody{
		Result: result,
		Count:  count,
	})
}

func GetProgConfig(c *gin.Context) {
	id, err := bindID(c)
	if err != nil {
		newClientError(c, err)
		return
	}
	progConfigService := service.NewProgConfigService()
	progConfig, err := progConfigService.GetProgConfigByID(id)
	if err != nil {
		newServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, vo.GetProgConfigResponseBody{
		Result: *progConfig,
	})
}

func UpdateProgConfig(c *gin.Context) {
	progConfigService := service.NewProgConfigService()
	profile := c.MustGet(PROFILE).(middleware.Profile)
	request := &vo.UpdateProgConfigRequest{}
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

	err = progConfigService.UpdateProgConfig(id, attrs)
	if err != nil {
		newServerError(c, err)
		return
	}
	newSuccess(c)
}

func DeleteProgConfig(c *gin.Context) {
	progConfigService := service.NewProgConfigService()
	id, err := bindID(c)
	if err != nil {
		newClientError(c, err)
		return
	}
	err = progConfigService.DeleteProgConfigByID(id)
	if err != nil {
		newServerError(c, err)
		return
	}
	newSuccess(c)
}
