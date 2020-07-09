package loginmiddleware

import (
	"net/http"

	"shadowDemo/shadow-framework/middleware"

	"github.com/gin-gonic/gin"
)

//TDefaultLoginHandler login middlewareHandler implementation
type TDefaultLoginHandler struct{}

func newDefaultLoginHandler() middleware.IMiddlewareHandler {
	return new(TDefaultLoginHandler)
}

//Handle handle redirect
func (handler *TDefaultLoginHandler) Handle(c *gin.Context) {
	Log.Debugln("loginHandler redirect to /")
	c.Redirect(http.StatusFound, "/")
}
