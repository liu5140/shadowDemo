package apimiddleware

import (
	"errors"
	"fmt"
	"net/http"
	"shadowDemo/APIServer/vo"
	"shadowDemo/model/do"
	"shadowDemo/service"
	"shadowDemo/shadow-framework/utils"

	"github.com/gin-gonic/gin"
)

func MerchantIPFilter(filterURLs ...string) gin.HandlerFunc {
	ipService := service.NewIPWhiteListService()

	return func(c *gin.Context) {
		if len(c.Errors) > 0 {
			return
		}
		if c.Request.Method == http.MethodPost {
			contain := false
			for _, u := range filterURLs {
				if c.Request.URL.RequestURI() == u {
					contain = true
					break
				}
			}
			if !contain {
				return
			}

			Log.Debug("MerchantIPFilter")
			ip := utils.GetRealIp(c.Request)

			if ip == "::1" {
				ip = "127.0.0.1"
			}
			Log.Debug("mcc login ip: ", ip)

			if len(c.Errors) > 0 {
				Log.Debug("ip filter has err")
				return
			}
			merchant, exist := c.Get(vo.PARAM_MERCHANT)
			if !exist {
				return
			}

			ipList, _ := ipService.GetIPWhiteListByAccountNo(merchant.(*do.Merchant).MerchantNo)

			//如果没有加入白名单数据，则白名单不起作用
			if len(ipList) == 0 {
				Log.Debug("no ip white list")
				return
			}

			Log.Debug("check ip white list")
			for _, IPWhite := range ipList {
				if ip == IPWhite.IP {
					return
				}
			}
			ipErr := fmt.Sprintf("IP White List filterd the IP %s of merchant %s ", ip, merchant.(*do.Merchant).Name)

			//	c.Error(IPNotAllowed{errors.New("ip not allowed")})
			c.Error(&gin.Error{
				Err:  errors.New(ipErr),
				Type: vo.IPVerifyFailed,
			})
		}
		c.Next()
	}
}
