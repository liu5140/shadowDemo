// Package router shadowDemo restful API.
//     Version: 1.0
//     Host: localhost:8030
//     BasePath: /acc
// SecurityDefinitions:
//     api_key:
//          type: apiKey
//          name: Authorization
//          in: header
// swagger:meta
package router

import (
	"errors"
	"net/http"
	"shadowDemo/middleware"
	"shadowDemo/model"
	"shadowDemo/service"
	"shadowDemo/zframework/middleware/concurrentlimit"
	"shadowDemo/zframework/middleware/ginlogrus"
	"shadowDemo/zframework/middleware/i18nmiddleware"
	"shadowDemo/zframework/middleware/securitymiddleware"
	modelc "shadowDemo/zframework/model"
	"shadowDemo/zframework/security"
	"shadowDemo/zframework/server"
	"shadowDemo/zframework/utils"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/i18n/language"
)

const (
	hmacSampleSecret = "tnb9Y0du$2a$10$KmatydruRTKlaUwErUOtNOXiPHVPunb9Y0dup9newm"
	hmacSecureSecret = "Tnb9Y0du$2a$10$KmatydruRTKlaUwErUOtNOXiPHVPunb9Y0dup9newm"
)

func MainRouter() http.Handler {
	loginPath := server.ServerConfigInstance().AdminServer.ContextPath + "/login"
	engine := gin.New()
	engine.Use(ginlogrus.Logger())
	engine.Use(gin.Recovery())
	engine.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "HEAD", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type, Content-Encoding, Authorization, CaptchaCode, CaptchaKey, OTP, Accept-Language, Accept-Encoding"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))
	//国际化处理
	engine.Use(i18nmiddleware.I18nResolver())
	//同iP 多少时间限制多少次访问
	engine.Use(concurrentlimit.ConcurrentLimit(100, 5*time.Minute))
	//jwt认证
	//engine.Use(middleware.JWTParse(hmacSampleSecret, 360000))
	//登录用户校验（冻结，密码错误）
	engine.Use(securitymiddleware.UsernamePasswordLoginFilter(loginPath))
	//获取当前用户信息
	//engine.Use(middleware.ProfileResolver(setAdmin))
	//谷歌验证码
	engine.Use(middleware.GoogleTokenValidator(loginPath, getProfileID))
	//普通验证码
	//	engine.Use(middleware.CaptchaHandler("/pcc/captcha", server.ServerConfigInstance().AdminServer.ContextPath))
	//验证权限
	//engine.Use(middleware.Authorizer(server.ServerConfig.ServerName))

	engine.Use(middleware.ErrorHandler())
	router := engine.Group(server.ServerConfigInstance().AdminServer.ContextPath)
	playerRouter(router)
	sessionRouter(router)
	webSocketRouter(router)
	monitorRouter(router)
	merchantRouter(router)
	progConfigRouter(router)
	aPIAccessReqLogRouter(router)
	aPIAccessResLogRouter(router)
	upmsMenuRouter(router)
	upmsRoleRouter(router)
	return engine
}

func getProfileID(c *gin.Context) (profileID int64, site string) {
	profile := c.MustGet(PROFILE).(middleware.Profile)
	site = profile.CurrentSite
	profileID = profile.ID
	return
}

func setAdmin(c *gin.Context) {
	token := c.MustGet(security.SHADOW_SECURITY_TOKEN).(*security.TUsernamePasswordAuthenticationToken)
	lang := c.MustGet("Lang").(*language.Language)
	adminService := service.NewUpmsAdminService()
	admin, err := adminService.GetUpmsAdminByLoginName(token.GetPrincipal())
	if err != nil {
		c.Error(err)
		return
	}
	ip := utils.GetRealIp(c.Request)
	profile := middleware.Profile{
		ID:          admin.ID,
		Username:    admin.RealName,
		Account:     admin.Account,
		UserType:    model.UserTypePlayer,
		Locked:      admin.State == modelc.Frozen,
		State:       admin.State,
		IP:          ip,
		DevInfo:     utils.GetDeviceInfo(c.Request),
		Lang:        lang.String(),
		Host:        server.ServerConfigInstance().AdminServer.Host,
		CurrentSite: server.ServerConfigInstance().AdminServer.CurrentSite,
	}
	c.Set(PROFILE, profile)

	if profile.Locked {
		c.Error(security.AccountLockedError{
			Err: errors.New("Account is locked"),
		})
	}
}
