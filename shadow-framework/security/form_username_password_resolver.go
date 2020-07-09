package security

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// TUsernamePasswordResolver default implementetion for UsernamePassworResolver
// It will resolver from form parameter
type TUsernamePasswordResolver struct{}

// NewFormUsernamePasswordResolver Constructor method for FormUsernaemPasswordResolver
func newFormUsernamePasswordResolver() IUsernamePasswordResolver {
	return &TUsernamePasswordResolver{}
}

func (formUsernamePasswordResolver TUsernamePasswordResolver) ObtainUsernamePassword(c *gin.Context) (string, string) {
	user := TUser{}
	if c.Bind(&user) == nil {
		return strings.TrimSpace(user.Username), user.Password
	}
	return "", ""
}
