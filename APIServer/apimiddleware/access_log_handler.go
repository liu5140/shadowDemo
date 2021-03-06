package apimiddleware

import (
	"encoding/json"
	"reflect"
	"shadowDemo/APIServer/vo"
	"shadowDemo/model/do"
	"shadowDemo/service"
	"shadowDemo/zframework/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func AccessLogHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.Errors) > 0 {
			return
		}
		orderParam, exist := c.Get(vo.PARAM_ORDER)
		var orderNo string
		if exist {
			order, _ := orderParam.(*do.Order)
			orderNo = order.OrderNo
		}
		req, _ := c.Get(vo.PARAM_REQUEST)
		mer, _ := c.Get(vo.PARAM_MERCHANT)
		merchant, _ := mer.(*do.Merchant)
		reqJSON, err := json.Marshal(&req)
		if err != nil {
			Log.Error(err)
		}

		reqLogService := service.NewAPIAccessReqLogService()

		urlPath := c.Request.URL.Path
		reqdatatype := reflect.TypeOf(req)
		var Request *vo.APIPublicRequest
		if reqdatatype.String() == "*vo.APIPublicRequest" {
			Request, _ = req.(*vo.APIPublicRequest)
			urlPath = urlPath + "(" + Request.Action + ")"
		}

		request := &do.APIAccessReqLog{
			MerchantName: merchant.Name,
			MerchantNo:   merchant.MerchantNo,
			IPAddress:    utils.GetRealIp(c.Request),
			APIURL:       urlPath,
			Request:      string(reqJSON),
			Method:       c.Request.Method,
			RequestTime:  time.Now(),
			OrderNo:      orderNo,
		}

		reqLogService.CreateAPIAccessReqLog(request)

		c.Next()
		//记录response
		res, _ := c.Get(vo.PARAM_RESPONSE)

		type APIPublicResponseLog struct {
			Code       int64
			ErrMsg     string
			Result     string
			ResultData interface{}
			Sign       string
		}

		datatype := reflect.TypeOf(res)
		resJSON := ""
		if datatype.String() == "vo.APIPublicResponse" {
			resConvert := res.(vo.APIPublicResponse)
			resLog := &APIPublicResponseLog{}
			resLog.Code = resConvert.Code
			resLog.ErrMsg = resConvert.ErrMsg
			resLog.Result = resConvert.Result
			resLog.ResultData = resConvert.ResultData
			resLog.Sign = resConvert.Sign

			resJ, err := json.Marshal(&resLog)
			if err != nil {
				Log.Error(err)
			}
			resJSON = string(resJ)
		} else {
			resJ, err := json.Marshal(&res)
			if err != nil {
				Log.Error(err)
			}
			resJSON = string(resJ)
		}

		resLogService := service.NewAPIAccessResLogService()

		resLogService.CreateAPIAccessResLog(&do.APIAccessResLog{
			RequestNo:    request.RequestNo,
			Response:     resJSON,
			ResponseTime: time.Now(),
			OrderNo:      orderNo,
		})
	}
}
