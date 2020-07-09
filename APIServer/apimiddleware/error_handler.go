package apimiddleware

import (
	"net/http"
	"net/url"
	"shadowDemo/shadow-framework/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

//ErrorHanlder 拦截请求参数校验错误， 如果有错误则直接返回，不记录请求日志log, 避免垃圾请求入库
//请求参数错误包括, 请求格式不能解析，请求token校验错误，商户不能解析
func ErrorHanlder() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.Errors) > 0 {
			err := c.Errors[0]
			Log.Error(err)
			errResp := ErrorResponse{}
			errResp.Code = int64(err.Type)
			errResp.ErrMsg = err.Error()
			errResp.Result = ""
			// 計算MD5
			//sec, _ := c.Get(bo.PARAM_SECRET)
			//secret, _ := sec.(*model.TBccSecret)
			//errResp = errResp.CalculateSign(secret.Secret)

			c.JSON(http.StatusOK, errResp)
			c.Abort()
			return
		}
	}
}

type ErrorResponse struct {
	Code   int64
	ErrMsg string
	Result string
	Sign   string
}

func (response ErrorResponse) CalculateSign(secret string) ErrorResponse {
	Log.Debug("CalculateSign Enter", secret)
	if secret == "" {
		return response
	}
	form := url.Values{}
	form.Add("Code", strconv.FormatInt(response.Code, 10))
	form.Add("Result", response.Result)
	enStr := form.Encode() + "&Key=" + secret
	Log.Debug("param: ", enStr)
	response.Sign = utils.MD5(enStr)
	return response
}
