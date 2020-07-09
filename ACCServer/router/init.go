package router

import (
	"errors"
	"net/http"
	"shadowDemo/ACCServer/router/vo"
	"shadowDemo/service"
	"shadowDemo/zframework/logger"
	"shadowDemo/zframework/security"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/nicksnyder/go-i18n/i18n"
)

var Log *logger.Logger

const (
	PROFILE        = "profile"
	MockPROFILE    = "mockprofile"
	SUCCESS        = "success"
	PARAM_SECRET   = "Secret"
	PARAM_MERCHANT = "Merchant"
	PARAM_REQUEST  = "Request"
	PARAM_PAYTYPE  = "PayType"
	PARAM_RESPONSE = "Response"
	PARAM_ORDER    = "Order"
	PARAM_ENDPOINT = "Endpoint"
	PARAM_ERROR    = "Error"
	PARAM_BODY     = "Body"
)

func init() {
	Log = logger.InitLog()
	Log.Infoln("注册表结构到数据库")
	//	model.InitialModels()
	Log.Infoln("登录方法注册")
	security.RegisterUserDetailService(security.USER_DETAILS_SERVICE, service.PlayerUserDetailService)
}

func bindID(c *gin.Context) (int64, error) {
	var idstr string
	idstr, ok := c.GetQuery("ID")
	if !ok {
		return 0, errors.New("ID is empty")
	}
	id64, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		Log.Error(err)
		return 0, err
	}
	return id64, nil
}

func newClientError(c *gin.Context, err error) {
	Log.Error(err)
	c.JSON(http.StatusBadRequest, vo.GenericMessageBody{
		Msg:  err.Error(),
		Code: http.StatusInternalServerError,
	})
}

func newServerError(c *gin.Context, err error) {
	Log.Error(err)
	T := c.MustGet("T").(i18n.TranslateFunc)
	message := ""
	switch err.(type) {
	case service.LessThanDepositError:
		message = T("key_deposit_notenough_error", map[string]interface{}{"limit": err.(service.LessThanDepositError).Limit})
	default:
		message = T("key_common_error")
	}

	if gorm.IsRecordNotFoundError(err) {
		c.JSON(http.StatusOK, vo.EmptyMessageBody{
			Result: make([]string, 0),
			Count:  0,
		})
	} else {
		c.JSON(http.StatusInternalServerError, vo.GenericMessageBody{
			Msg:  message,
			Code: http.StatusBadRequest,
		})
	}
}

func newSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, vo.GenericMessageBody{
		Msg:  "success",
		Code: http.StatusOK,
	})
}
