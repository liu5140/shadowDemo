package apimiddleware

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"reflect"
	"shadowDemo/APIServer/vo"
	"shadowDemo/service"

	"shadowDemo/shadow-framework/utils/aes"

	"github.com/gin-gonic/gin"
)

func RequestResolver() gin.HandlerFunc {
	return func(c *gin.Context) {
		Log.Debug("RequestResolver")
		Log.Debug("request length: ", c.Request.ContentLength)
		if c.Request.ContentLength >= 10*1024 {
			Log.Error("Content Length is too long")
			c.Error(errors.New("Content Length is too long"))
			return
		}

		var request interface{}
		switch c.Request.URL.Path {
		case ACTION_URL:
			request = &vo.APIPublicRequest{}
		case VIEW_DEPOSIT_URL:
			request = &vo.APIViewDepositRequest{}
		default:
			Log.WithField("path", c.Request.URL.Path).Error("invalid request URL")
			c.Abort()
			return
		}

		err := c.Bind(request)
		if err != nil {
			c.Error(err)
			return
		}
		reqJSON, _ := json.Marshal(request)
		Log.Debug("checkValidText reqJSON: ", string(reqJSON))

		c.Set(vo.PARAM_REQUEST, request)
		datatype := reflect.TypeOf(request)
		if datatype.String() == "*vo.APIPublicRequest" && request.(*vo.APIPublicRequest).Action != QUERY_BALANCE {
			err = bindingData(c, request.(*vo.APIPublicRequest))
			if err != nil {
				Log.Error(err)
				Log.Debug("Request Resolve failed")
				c.Error(err)
			}
		}
	}
}

func bindingData(c *gin.Context, req *vo.APIPublicRequest) error {
	Log.Debug("bindingData, action: ", req.Action)
	//通过商户资料获取商户的基本信息
	merchantService := service.NewMerchantService()
	merchant, err := merchantService.GetMerchantByMerNo(req.MerchantNo)
	if err != nil {
		Log.Debug("Get Merchant Key failed")
		return errors.New("Get Merchant Key failed")
	}
	//把商户信息放到c中
	c.Set(vo.PARAM_MERCHANT, merchant)
	Log.Debug("AesKey: ", merchant.AesKey)
	aeskey := merchant.AesKey
	//把字符串从HEX转成byte数组
	deHexstr, err := hex.DecodeString(req.Data)
	if err != nil {
		Log.Error(err)
		return err
	}
	Log.Infoln("========", deHexstr)
	Log.Debug("plain: ", len(deHexstr))
	//然后进行aes解密
	decryptstr, err := aes.AesEcbDecrypt([]byte(deHexstr), []byte(aeskey))
	if err != nil {
		Log.Error(err)
		return err
	}
	//获取解密后的字符串
	decryptData := string(decryptstr)
	Log.Debug("decryptData: ", decryptData)

	//获取后把字符串序列化为结构体
	switch req.Action {
	case BANK_LIST:
		var data vo.APIChannelReqData
		err := req.UnserializeFromJson(decryptData, &data)
		if err != nil { // 解析失败则清理缓存重新加载
			Log.Errorf("parse channel request data field, data=%v, err=%v", req.Data, err)
			return err
		}
		req.DataStruct = data

	case DEPOSIT:
		var data vo.APIDepositReqData
		err := req.UnserializeFromJson(decryptData, &data)
		if err != nil { // 解析失败则清理缓存重新加载
			Log.Errorf("parse deposit deposit data field, data=%v, err=%v", req.Data, err)
			return err
		}
		req.DataStruct = data

	case WITHDRAW:
		var data vo.APIWithdrawReqData
		err := req.UnserializeFromJson(decryptData, &data)
		if err != nil { // 解析失败则清理缓存重新加载
			Log.Errorf("parse withdraw request data field, data=%v, err=%v", req.Data, err)
			return err
		}
		req.DataStruct = data

	case QUERY_DEPOSIT:
		var data vo.APIQueryDepositReqData
		err := req.UnserializeFromJson(decryptData, &data)
		if err != nil { // 解析失败则清理缓存重新加载
			Log.Errorf("parse query order request data field, data=%v, err=%v", req.Data, err)
			return err
		}
		req.DataStruct = data
	case QUERY_BALANCE:
		var data vo.APIQueryBalanceReqData
		err := req.UnserializeFromJson(decryptData, &data)
		if err != nil { // 解析失败则清理缓存重新加载
			Log.Errorf("parse query order request data field, data=%v, err=%v", req.Data, err)
			return err
		}
		req.DataStruct = data
	}
	check := req.DataStruct.(IDataResolveCheck)
	err = check.CheckRequestResolve()
	if err != nil {
		return err
	}
	return nil
}

type IDataResolveCheck interface {
	CheckRequestResolve() error
}
