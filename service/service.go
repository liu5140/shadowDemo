package service

type Service struct {
	APIService             *APIService
	APIAccessReqLogService *APIAccessReqLogService
	APIAccessResLogService *APIAccessResLogService
	GoogleTokenService     *GoogleTokenService
	IPWhiteListService     *IPWhiteListService
	MerchantService        *MerchantService
	OrderService           *OrderService
	PlayerService          *PlayerService
	ProgConfigService      *ProgConfigService
	SecureConfigureService *SecureConfigureService
	SessionService         *SessionService
	UserService            *UserService
}

var service *Service = nil

func ServiceInit() {
	service = &Service{}
	service.APIService = NewAPIService()
	service.APIAccessReqLogService = NewAPIAccessReqLogService()
	service.APIAccessResLogService = NewAPIAccessResLogService()
	service.GoogleTokenService = NewGoogleTokenService()
	service.IPWhiteListService = NewIPWhiteListService()
	service.MerchantService = NewMerchantService()
	service.OrderService = NewOrderService()
	service.PlayerService = NewPlayerService()
	service.ProgConfigService = NewProgConfigService()
	service.SecureConfigureService = NewSecureConfigureService()
	service.SessionService = NewSessionService()
	service.UserService = NewUserService()
}
