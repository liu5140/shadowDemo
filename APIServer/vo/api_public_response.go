package vo

import (
	"encoding/hex"
	"net/url"
	"shadowDemo/shadow-framework/utils"
	"shadowDemo/shadow-framework/utils/aes"
	"shadowDemo/shadow-framework/utils/encoder"
	"strconv"
	"strings"
)

type APIPublicResponse struct {
	Code       int64
	ErrMsg     string `url:"-"`
	Result     string
	ResultData interface{} `json:"-"`
	Sign       string
}

//结果加密
func (response APIPublicResponse) EncryptData(enckey string) (APIPublicResponse, error) {
	plain, err := encoder.JSONMarshal(response.ResultData)
	if err != nil {
		return response, err
	}
	if response.ResultData == nil {
		return response, err
	}
	Log.Debug("plain: ", string(plain))
	Log.Debug("key: ", enckey)
	mm, err := aes.AesEcbEncrypt(plain, []byte(enckey))
	if err != nil {
		return response, err
	}
	mmstr := hex.EncodeToString(mm)
	Log.Debug("EncryptData: ", mmstr)
	response.Result = strings.ToUpper(string(mmstr))
	return response, nil
}

//返回结果计算sign
func (response APIPublicResponse) CalculateSign(secret string) APIPublicResponse {
	Log.Debug("CalculateSign Enter", secret)
	form := url.Values{}
	form.Add("Code", strconv.FormatInt(response.Code, 10))
	form.Add("Result", response.Result)
	enStr := form.Encode() + "&Key=" + secret
	Log.Debug("param: ", enStr)
	response.Sign = utils.MD5(enStr)
	return response
}

//验证sign
func (response *APIPublicResponse) VerifySign(secret string) bool {
	Log.Debug("response.VerifySign secret: ", secret)
	//计算签名
	sign := response.CalculateSign(secret).Sign
	Log.Debug("response.VerifySign.CalculateSign sign: ", sign)
	Log.Debug("reponse sign: ", response.Sign)
	return strings.ToUpper(sign) == strings.ToUpper(response.Sign)
}

//解密
func (response *APIPublicResponse) DecryptData(aesKey string) (data []byte, err error) {
	Log.Debug("response.DecryptData AesKey: ", aesKey)
	deHexstr, err := hex.DecodeString(response.Result)
	if err != nil {
		return nil, err
	}
	data, err = aes.AesEcbDecrypt([]byte(deHexstr), []byte(aesKey))
	if err != nil {
		return nil, err
	}

	return data, nil
}
