package router

import (
	"shadowDemo/ACCServer/router/vo"
	"shadowDemo/model/do"
	"shadowDemo/service"
	"strings"

	"shadowDemo/zframework/utils"

	"github.com/gin-gonic/gin"
)

func merchantRouter(r *gin.RouterGroup) {

	r.POST("/merchant", createMerchant)

}

func createMerchant(c *gin.Context) {
	merService := service.NewMerchantService()
	request := &vo.CreateMerchantRequest{}
	err := c.Bind(&request.Body)
	if err != nil {
		newClientError(c, err)
		return
	}
	Log.Debug("createMerchant request: ", utils.FormatStruct(request))

	mer := do.Merchant{
		MerchantNo:    strings.TrimSpace(request.Body.MerchantNo),
		LoginPassword: strings.TrimSpace(request.Body.LoginPassword),
	}

	err = merService.CreateMerchant(&mer)
	if err != nil {
		newServerError(c, err)
		return
	}
	newSuccess(c)
}
