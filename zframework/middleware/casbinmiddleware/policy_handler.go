package casbinmiddleware

import (
	"net/http"
	"strings"

	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
)

//TODO add a entry api for update policy

// Policy use to get, save , modify, delete policy
func Policy(enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.Request.URL.RequestURI()
		method := c.Request.Method
		if url == "/policy" {

			switch method {
			case http.MethodGet:
				doGetPolicy(c, enforcer)
			case http.MethodPost:
				doSavePolicy(c, enforcer)
			case http.MethodPut:
				doUpdatePolicy(c, enforcer)
			case http.MethodDelete:
				doDeletePolicy(c, enforcer)
			}
		}

		if url == "/role" {
			switch method {
			case http.MethodGet:
			case http.MethodPost:
			case http.MethodPut:
			case http.MethodDelete:
			}
		}

	}
}

func doGetPolicy(c *gin.Context, enforcer *casbin.Enforcer) {
	c.Set("actions", enforcer.GetAllActions())
}

func doSavePolicy(c *gin.Context, enforcer *casbin.Enforcer) {
	sub := c.GetString("sub")
	domain := c.GetString("domain")
	obj := c.GetString("obj")
	act := c.GetString("act")
	if strings.TrimSpace(domain) != "" {
		enforcer.Enforce(sub, domain, obj, act)
	} else {
		enforcer.Enforce(sub, obj, act)
	}
	enforcer.SavePolicy()

}
func doDeletePolicy(c *gin.Context, enforcer *casbin.Enforcer) {

}
func doUpdatePolicy(c *gin.Context, enforcer *casbin.Enforcer) {

}

func doDeleteRole(c *gin.Context, enforcer *casbin.Enforcer) {

}
