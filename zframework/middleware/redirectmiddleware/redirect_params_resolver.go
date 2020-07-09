package redirectmiddleware

import (
	"encoding/gob"
	"net/http"

	"shadowDemo/zframework/middleware/sessionmiddleware"

	"github.com/gin-gonic/gin"
)

func RedirectParamsResolver() gin.HandlerFunc {
	return func(c *gin.Context) {
		sess := sessionmiddleware.GetCurrentSession(c)
		param := sess.Get(REDIRECT_PARAMETER_KEY)

		if param != nil {
			for key, val := range param.(map[string]interface{}) {
				c.Set(key, val)
			}
			sess.Delete(REDIRECT_PARAMETER_KEY)
		}
		c.Next()
	}
}

func RedirectWithParams(c *gin.Context, url string, params map[string]interface{}) {
	sess := sessionmiddleware.GetCurrentSession(c)
	sess.Set(REDIRECT_PARAMETER_KEY, params)
	for _, val := range params {
		gob.Register(val)
	}
	c.Redirect(http.StatusSeeOther, url)
}
