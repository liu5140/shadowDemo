package middleware

import (
	"errors"
	"net/http"
	"strconv"

	"shadowDemo/service"

	"shadowDemo/zframework/security"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func JWTParse(secret string, maxLifeTime int64) gin.HandlerFunc {
	remitList := map[string]string{
		"/acc/server":   http.MethodGet,
		"/acc/merchant": http.MethodPost,
	}

	return func(c *gin.Context) {
		Log.Debug("jwt processing")
		if len(c.Errors) > 0 {
			return
		}
		if v, ok := remitList[c.Request.URL.Path]; ok && c.Request.Method == v {
			Log.Debug("ok && c.Request.Method == v")
			anonymousToken := security.NewAnonymousAuthenticationToken()
			anonymousToken.SetDetails(&security.TWebAuthenticationDetails{
				RemoteAddress: c.ClientIP(),
				RequestURI:    c.Request.URL.Path,
			})
			c.Set(security.SHADOW_SECURITY_TOKEN, anonymousToken)
			return
		}

		// 获取token
		token, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
			Log.Debug("jwt token")
			b := ([]byte(secret))
			return b, nil
		})

		if err != nil {
			Log.Error("jwt token error", err)
			anonymousToken := security.NewAnonymousAuthenticationToken()
			anonymousToken.SetDetails(&security.TWebAuthenticationDetails{
				RemoteAddress: c.ClientIP(),
				RequestURI:    c.Request.URL.Path,
			})
			c.Set(security.SHADOW_SECURITY_TOKEN, anonymousToken)
			err = security.JwtExpiredErr{
				Err: errors.New("jwt 失效2"),
			}
			c.Error(err)
			//c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		// 校验并解析token
		var account, profileID string
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			Log.WithFields(logrus.Fields{
				"PID":     claims["PID"],
				"Account": claims["Account"],
				"Time":    claims["Time"],
			}).Debug("jwt parse")
			profileID = claims["PID"].(string)
			account = claims["Account"].(string)
			id64, _ := strconv.ParseInt(profileID, 10, 64)
			profile := Profile{
				ID:      id64,
				Account: account,
			}
			c.Set("mockprofile", profile)
		} else {
			Log.Debug("security.SHADOW_SECURITY_TOKEN")
			anonymousToken := security.NewAnonymousAuthenticationToken()
			anonymousToken.SetDetails(&security.TWebAuthenticationDetails{
				RemoteAddress: c.ClientIP(),
				RequestURI:    c.Request.URL.Path,
			})
			c.Set(security.SHADOW_SECURITY_TOKEN, anonymousToken)
			return
		}

		// 验证session是否过期
		sessionService := service.NewSessionService()
		err = sessionService.SetSessionExpireTime(profileID, token.Raw, maxLifeTime)
		if err != nil {
			Log.Error("validate session expired :", err)
			err = security.JwtExpiredErr{
				Err: errors.New("jwt 失效2"),
			}
			c.Error(err)
			//c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		// 创建权限认证token
		authRequest := security.NewUsernamePasswordAuthenticationToken(account, profileID)
		authRequest.SetAuthenticated(true)
		c.Set(security.SHADOW_SECURITY_TOKEN, authRequest)
		c.Set("jwt_token", token.Raw)
		if err != nil {
			Log.Error(err)
			//	c.AbortWithError(http.StatusUnauthorized, err)
			err = security.JwtExpiredErr{
				Err: errors.New("jwt 失效"),
			}
			c.Error(err)
			return
		}
	}
}
