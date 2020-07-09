package router

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"shadowDemo/model"
	"strings"

	"shadowDemo/ACCServer/router/vo"
	"shadowDemo/middleware"
	"shadowDemo/model/do"
	"shadowDemo/service"
	modelc "shadowDemo/shadow-framework/model"
	"shadowDemo/shadow-framework/server"

	"shadowDemo/shadow-framework/utils"

	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/i18n/language"
)

func sessionRouter(r *gin.RouterGroup) {
	// swagger:route POST /logout session logout
	//
	// 玩家登出系统;
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
	r.POST("/logout", logout)
	// swagger:route POST /login session login
	//
	// 玩家登录系统;
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http
	//
	//	     Security:
	//       api_key:
	//       oauth: read, write
	//
	//     Responses:
	//       200: loginResponse
	//       default: genericError
	r.POST("/login", login)
	// swagger:route POST /sms/logins session smslogin
	//
	// 玩家短信登录系统;
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http
	//
	//	     Security:
	//       api_key:
	//       oauth: read, write
	//
	//     Responses:
	//       200: loginResponse
	//       default: genericError
	r.POST("/sms/login", smslogin)
	// swagger:route GET /mock/token session mockToken
	//
	// 未登陆获取一个默认token;
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http
	//
	//     Responses:
	//       200: loginResponse
	//       default: genericError
	r.GET("/mock/token", mockToken)
}

func logout(c *gin.Context) {
	profile := c.MustGet(PROFILE).(middleware.Profile)
	tokenString := c.MustGet("jwt_token").(string)
	sessionService := service.NewSessionService()

	//logout已进入的游戏
	hash, err := sessionService.GetAllSessionContent(fmt.Sprint(profile.ID))
	if err != nil {
		newServerError(c, err)
		return
	}

	var endpoints []string
	for k, v := range hash {
		if strings.HasPrefix(k, "login:endpoint:key:") {
			endpoints = append(endpoints, v)
		}
	}

	//logout session
	sessionService.DeleteSessionByToken(tokenString)

	c.JSON(http.StatusOK, "success")
}

func login(c *gin.Context) {
	profile := c.MustGet(PROFILE).(middleware.Profile)
	mockProfile := c.MustGet(MockPROFILE).(middleware.Profile)
	Log.Infoln("======", utils.FormatStruct(mockProfile))
	playerService := service.NewPlayerService()
	var err error
	player, err := playerService.GetPlayerByID(profile.ID)
	if err != nil {
		newServerError(c, err)
		return
	}
	LoginSer(c, player)
}

func smslogin(c *gin.Context) {
	playerService := service.NewPlayerService()
	request := &vo.SmsAccountRequest{}
	err := c.Bind(request)
	if err != nil {
		newClientError(c, err)
		return
	}
	player, err := playerService.GetPlayerByLoginName(request.UserName)
	if err != nil {
		newServerError(c, err)
		return
	}
	setPlayerAuto(c, &player)
	LoginSer(c, player)
}

func LoginSer(c *gin.Context, player do.Player) {
	playerService := service.NewPlayerService()
	sessionService := service.NewSessionService()
	profile := c.MustGet(PROFILE).(middleware.Profile)
	u, err := url.Parse(c.GetHeader("referer"))
	if err != nil {
		newServerError(c, err)
		return
	}

	//更新用户最后登录时间
	err = playerService.UpdateLastLoginTime(profile.ID, profile.IP, profile.IPAddr, u.Host, profile.DevInfo)
	if err != nil {
		newServerError(c, err)
		return
	}

	//创建token
	tokenString, err := sessionService.CreateJWT(fmt.Sprintf("%d", profile.ID), profile.Account, hmacSampleSecret)
	if err != nil {
		newServerError(c, err)
		return
	}

	err = sessionService.CreateSession(profile.ID, profile.UserType, profile.IP, profile.DevInfo, "ss", tokenString)
	if err != nil {
		newServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, vo.LoginResponseBody{
		Token:    tokenString,
		ID:       profile.ID,
		Account:  player.Account,
		NickName: player.NickName,
		Mobile:   player.Mobile,
		Birthday: player.Birthday,
		//Sex:            player.Sex,
		Country: player.Country,
		//Balance:        player.Wallet.Balance,
		HostIndex:      0,
		Recommend:      player.Recommend,
		IsSecurePwdSet: player.IsSecurePwdSet,
		Code:           http.StatusOK,
	})
}

func setPlayerAuto(c *gin.Context, player *do.Player) (err error) {
	Log.Debug("setPlayer:", player)
	lang := c.MustGet("Lang").(*language.Language)
	ip := utils.GetRealIp(c.Request)
	profile := middleware.Profile{
		ID:       player.ID,
		Username: player.NickName,
		Account:  player.Account,
		UserType: model.UserTypePlayer,
		Locked:   player.State == modelc.Frozen,
		State:    player.State,
		IP:       ip,
		DevInfo:  utils.GetDeviceInfo(c.Request),
		Lang:     lang.String(),
		Host:     server.ServerConfigInstance().PlayerServer.Host,
	}
	c.Set(PROFILE, profile)
	if profile.Locked {
		err = errors.New("Account is locked")
		return
	}
	return
}

//生成一个测试得账号
func mockToken(c *gin.Context) {
	sessionService := service.NewSessionService()
	//获取一个用户id
	playerID := service.GenPlayerID()
	pAccount := fmt.Sprintf("mock_%d", playerID)

	tokenString, err := sessionService.CreateJWT(fmt.Sprintf("%d", playerID), pAccount, hmacSampleSecret)
	if err != nil {
		newServerError(c, err)
		return
	}

	err = sessionService.CreateSession(playerID, model.UserTypeMockPlayer, utils.GetRealIp(c.Request), utils.GetDeviceInfo(c.Request), "ss", tokenString)
	if err != nil {
		newServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, vo.LoginResponseBody{
		Token:   tokenString,
		ID:      playerID,
		Account: pAccount,
		Code:    http.StatusOK,
	})
}
