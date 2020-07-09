package vo

import (
	"encoding/hex"
	"encoding/json"
	"net/url"
	"shadowDemo/zframework/utils"
	"shadowDemo/zframework/utils/aes"
	"shadowDemo/zframework/utils/encoder"
	"shadowDemo/zframework/utils/query"
	"strings"
)

type APIPublicRequest struct {
	MerchantNo string      `binding:"required"`
	Action     string      `binding:"required"`
	Data       string      `binding:"required"`
	DataStruct interface{} `url:"-"`
	Sign       string      `binding:"required" url:"-"`
}

func (request *APIPublicRequest) EncryptData(enckey string) error {
	plain, err := encoder.JSONMarshal(request.DataStruct)
	Log.Debug("push message EncryptData request.DataStruct:", plain)
	if err != nil {
		return err
	}
	mm, err := aes.AesEcbEncrypt(plain, []byte(enckey))
	if err != nil {
		return err
	}
	mmstr := hex.EncodeToString(mm)
	request.Data = strings.ToUpper(string(mmstr))
	Log.Debug("EncryptData Data:", request.Data)

	return nil
}

func (request APIPublicRequest) VerifySign(secret string) bool {
	values, _, _ := query.Values(request)
	param := values.Encode() + "&Key=" + secret
	Log.Debug("Key: ", secret)
	Log.Debug("param: ", param)
	sign := strings.ToUpper(encoder.MD5(param))
	Log.Debug("sign: ", sign)
	Log.Debug("request sign: ", request.Sign)
	return sign == strings.ToUpper(request.Sign)
}

func (request *APIPublicRequest) CalculateSign(secret string) {
	Log.Debug("CalculateSign Enter with the secret: ", secret)
	form := url.Values{}
	form.Add("Action", request.Action)
	form.Add("Data", request.Data)
	form.Add("MerchantNo", request.MerchantNo)
	enStr := form.Encode() + "&Key=" + secret
	Log.Debug("md5 with param:", enStr)
	request.Sign = utils.MD5(enStr)
	Log.Debug("callback sign:", request.Sign)
}

func (request APIPublicRequest) GetMerchantNo() string {
	return request.MerchantNo
}

func (data APIPublicRequest) UnserializeFromJson(jsonstr string, st interface{}) error {
	d := json.NewDecoder(strings.NewReader(jsonstr))
	d.UseNumber()
	return d.Decode(st)
}
