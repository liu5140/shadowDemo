package sessionmiddleware

import (
	"github.com/astaxie/beego/session"
	"github.com/gin-gonic/gin"
)

const (
	gSessionID = "gessionid"
)

// Sessions session middleware
func Sessions(globalSessions *session.Manager) gin.HandlerFunc {
	if globalSessions == nil {
		panic("globalSessions is nil")
	}
	return func(c *gin.Context) {
		sess, err := globalSessions.SessionStart(c.Writer, c.Request)
		if err != nil {
			Log.Panic(err)
		}
		defer sess.SessionRelease(c.Writer)
		c.Set(gSessionID, sess)
		c.Next()
	}
}

// GetCurrentSession get sesison from current request
func GetCurrentSession(c *gin.Context) session.Store {
	return c.MustGet(gSessionID).(session.Store)
}
