package router

import (
	"net/url"
	"shadowDemo/ACCServer/router/vo"
	"shadowDemo/middleware"
	"shadowDemo/model/do"
	"shadowDemo/service"
	"strings"

	"shadowDemo/shadow-framework/utils"

	"github.com/gin-gonic/gin"
)

func playerRouter(r *gin.RouterGroup) {
	// swagger:route POST /player player createPlayer
	//
	// 创建玩家;
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
	r.POST("/player", createPlayer)

}

func createPlayer(c *gin.Context) {
	playerService := service.NewPlayerService()
	request := &vo.CreatePlayerRequest{}
	err := c.Bind(&request.Body)
	if err != nil {
		newClientError(c, err)
		return
	}

	player := do.Player{
		Account:          strings.TrimSpace(request.Body.Account),
		PhoneCode:        strings.TrimSpace(request.Body.PhoneCode),
		LoginPassword:    strings.TrimSpace(request.Body.LoginPassword),
		RecommanderPCode: strings.TrimSpace(request.Body.Recommend),
	}

	ip := utils.GetRealIp(c.Request)
	u, err := url.Parse(c.GetHeader("referer"))
	if err != nil {
		newServerError(c, err)
		return
	}
	player.RegistURL = u.Host
	player.RegistIP = ip
	//player.CreateDevice = utils.GetDeviceInfo(c.Request)
	Log.Debug("pccCreatePlayer request: ", utils.FormatStruct(request))
	err = playerService.CreatePlayer(0, &player)
	if err != nil {
		newServerError(c, err)
		return
	}
	//此时获取得是mock账户得id 和账号
	profile := c.MustGet(PROFILE).(middleware.Profile)
	Log.Infoln("=========", utils.FormatStruct(profile))
	//合并订单 把mock得id合并到 注册用户得id

	newSuccess(c)
}
