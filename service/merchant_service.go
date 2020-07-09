package service

import (
	"shadowDemo/model"
	"shadowDemo/model/do"

	shadowsecurity "shadowDemo/zframework/security"
	"shadowDemo/zframework/utils"
	"shadowDemo/zframework/utils/encrypt"
)

type MerchantService struct {
	passwordEncoder shadowsecurity.IPasswordEncoder
	encryptKey      []byte
	encryptKey2     []byte
}

var merchantService *MerchantService

func NewMerchantService() *MerchantService {
	if merchantService == nil {
		l.Lock()
		if merchantService == nil {
			merchantService = &MerchantService{
				passwordEncoder: shadowsecurity.PasswordEncoderInstance(shadowsecurity.PASSWORD_ENCODER),
				encryptKey:      []byte("g5sbUQIO6IVpvJSINIQd3G8oTBORzClg"),
			}
		}
		l.Unlock()
	}
	return merchantService
}

func (service *MerchantService) CreateMerchant(mer *do.Merchant) error {
	mer.AesKey = encrypt.Encrypt(service.encryptKey, utils.RandString(8))
	mer.Md5Secret = encrypt.Encrypt(service.encryptKey, utils.RandString(16))
	mer.LoginPassword = service.passwordEncoder.Encode(mer.LoginPassword)
	return model.GetModel().MerchantDao.Create(mer)
}

func (service *MerchantService) GetMerchantByMerNo(merNo string) (*do.Merchant, error) {
	m := model.GetModel()
	result := &do.Merchant{
		MerchantNo: merNo,
	}
	err := m.MerchantDao.FindOne(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
