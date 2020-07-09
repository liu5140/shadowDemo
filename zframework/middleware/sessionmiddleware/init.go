package sessionmiddleware

import (
	"encoding/gob"

	"shadowDemo/zframework/logger"
	"shadowDemo/zframework/security"
)

var (
	Log *logger.Logger
)

func init() {
	Log = logger.InitLog()
	Log.Info("session middleware init")
	gob.Register(&security.TAnonymousAuthenticationToken{})
	gob.Register(&security.TRequestAuthenticationToken{})
	gob.Register(&security.TUsernamePasswordAuthenticationToken{})
	gob.Register(&security.TWebAuthenticationDetails{})
	gob.Register(&security.TUser{})
}
