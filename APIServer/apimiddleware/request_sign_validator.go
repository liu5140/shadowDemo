package apimiddleware

import (
	"errors"
	"shadowDemo/APIServer/vo"
	"shadowDemo/model/do"

	"github.com/gin-gonic/gin"
)

func RequestSignValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.Errors) > 0 {
			return
		}
		req, _ := c.Get(vo.PARAM_REQUEST)       //request
		merchant, _ := c.Get(vo.PARAM_MERCHANT) //商户信息
		verifyReq, _ := req.(IVerifiedRequest)
		mer, _ := merchant.(*do.Merchant)
		if !verifyReq.VerifySign(mer.Md5Secret) {
			c.Error(errors.New("Sign validate failed"))
		}
	}
}

type IVerifiedRequest interface {
	VerifySign(secret string) bool
}
