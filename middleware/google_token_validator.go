package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"shadowDemo/service"

	"github.com/gin-gonic/gin"
)

func GoogleTokenValidator(loginPath string, fn func(c *gin.Context) (profileID int64, site string)) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.Errors) > 0 {
			return
		}

		if c.Request.URL.RequestURI() == loginPath && c.Request.Method == http.MethodPost {
			profileID, site := fn(c)
			passCode := c.GetHeader("OTP")
			if passCode == "" {
				c.Error(WrongGoogleToken{
					Err: errors.New("google token is empty"),
				})
				return
			}
			googleTokenService := service.NewGoogleTokenService()
			if ok := googleTokenService.VerifyTotp(site, fmt.Sprintf("%d", profileID), passCode); !ok {
				c.Error(WrongGoogleToken{
					Err: errors.New("google token is wrong"),
				})
			}
		}
	}
}
