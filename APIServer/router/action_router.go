package router

import (
	"net/http"
	"shadowDemo/APIServer/apimiddleware"
	"shadowDemo/APIServer/vo"
	"shadowDemo/model/do"
	"shadowDemo/service"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

func actionRouter(r *gin.RouterGroup) {
	r.POST("/action", action)
}

func action(c *gin.Context) {
	Log.Debug("/action Enter")
	req, _ := c.Get(vo.PARAM_REQUEST)
	request, _ := req.(*vo.APIPublicRequest)
	middleErr, _ := c.Get(vo.PARAM_ERROR)
	merchant, _ := c.Get(vo.PARAM_MERCHANT) //商户信息
	mer, _ := merchant.(*do.Merchant)
	var resp vo.APIPublicResponse
	apiservice := service.NewAPIService()
	Log.Debug("request: ", request)

	switch request.Action {
	case apimiddleware.BANK_LIST:
		result, err := apiservice.DoGetBanklist(c, request)
		if err == nil && result != nil {
			resp.Code = 0
			resp.ErrMsg = ""
			resp.ResultData = result.Data
		} else {
			resp.Code = -1
			resp.ErrMsg = err.Error()
		}
	case apimiddleware.DEPOSIT:
		if middleErr != nil {
			resp.Code = -1 //int64(vo.OrderCreateFailed)
			resp.ErrMsg = middleErr.(string)
		} else {
			result, err := apiservice.DoDeposit(c, request)
			if err == nil {
				resp.Code = 0
				resp.ErrMsg = ""
				resp.ResultData = result
			} else {
				resp.Code = -1
				resp.ErrMsg = err.Error()
			}
		}
	case apimiddleware.WITHDRAW:
		if middleErr != nil {
			resp.Code = -1
			resp.ErrMsg = middleErr.(string)
		} else {
			result, err := apiservice.DoWithdraw(c, request)
			if err == nil {
				resp.Code = 0
				resp.ErrMsg = ""
				resp.ResultData = result
			} else {
				resp.Code = -1
				resp.ErrMsg = err.Error()
			}
		}
	case apimiddleware.QUERY_DEPOSIT:
		if middleErr != nil {
			resp.Code = -1 //int64(bo.QueryOrderFailed)
			resp.ErrMsg = middleErr.(string)
		} else {
			result, err := apiservice.DoQueryOrder(c, request)
			if err == nil {
				resp.Code = 0
				resp.ErrMsg = ""
				resp.ResultData = result
			} else {
				resp.Code = -1
				resp.ErrMsg = err.Error()
			}
		}
	case apimiddleware.QUERY_BALANCE:
		if middleErr != nil {
			resp.Code = -1 //int64(bo.RequestResolveFailed)
			resp.ErrMsg = middleErr.(string)
		} else {
			//	merchant := c.MustGet(vo.PARAM_MERCHANT).(*do.Merchant)
			//	fundService := service.NewBccFundChangeService()
			//balance, frozenBalance := fundService.GetMerchantBalanceAndFrozenBalance(merchant.ID)
			resp.Code = 0
			resp.ErrMsg = ""
			resp.ResultData = vo.APIQueryBalanceResult{
				Balance:       decimal.Zero,
				FrozenBalance: decimal.Zero,
			}
		}
	default:
		resp.Code = -1 //int64(vo.RequestResolveFailed)
		resp.ErrMsg = "Request resolve failed"
	}

	Log.Debug("resp: ", resp)

	resp, err := resp.EncryptData(mer.AesKey)
	if err != nil {
		resp.Result = ""
		resp.ResultData = nil
		c.JSON(http.StatusBadRequest, resp)
		c.Set(vo.PARAM_RESPONSE, resp)
		Log.Debug("/action End")
		return
	}
	// 計算MD5
	resp = resp.CalculateSign(mer.Md5Secret)

	c.JSON(http.StatusOK, resp)
	c.Set(vo.PARAM_RESPONSE, resp)
	Log.Debug("/action End")
}
