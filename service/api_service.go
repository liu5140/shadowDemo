package service

import (
	"shadowDemo/APIServer/vo"

	"github.com/gin-gonic/gin"
)

type APIService struct {
}

var apiService *APIService

func NewAPIService() *APIService {
	if apiService == nil {
		l.Lock()
		if apiService == nil {
			apiService = &APIService{}
		}
		l.Unlock()
	}
	return apiService
}

func (service *APIService) DoGetBanklist(c *gin.Context, request *vo.APIPublicRequest) (*vo.APIChannelResult, error) {
	var datas []vo.APIChannelResData
	for i := 0; i < 10; i++ {
		var data vo.APIChannelResData
		data.BankCode = "13213232"
		datas = append(datas, data)
	}
	return &vo.APIChannelResult{
		Data: datas,
	}, nil
}

func (service *APIService) DoDeposit(c *gin.Context, request *vo.APIPublicRequest) (*vo.APIDepositResult, error) {
	return nil, nil
}

func (service *APIService) DoWithdraw(c *gin.Context, request *vo.APIPublicRequest) (*vo.APIWithdrawResult, error) {
	return nil, nil
}

func (service *APIService) DoQueryOrder(c *gin.Context, request *vo.APIPublicRequest) (*vo.APIQueryDepositResult, error) {
	return nil, nil
}
