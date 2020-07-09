package validator

import (
	"shadowDemo/shadow-framework/logger"

	"github.com/gin-gonic/gin/binding"
)

var Log *logger.Logger

func init() {
	Log = logger.InitLog()
	Log.Infoln("Validator init")

}

func init() {
	binding.Validator = new(defaultValidator)
}
