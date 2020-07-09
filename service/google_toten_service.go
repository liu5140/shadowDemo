package service

import (
	"bytes"
	"image/png"

	"github.com/pquerna/otp/totp"
)

type GoogleTokenService struct {
}

var googleTokenService *GoogleTokenService

const accountPrefix = "google:secret:"

func NewGoogleTokenService() *GoogleTokenService {
	if googleTokenService == nil {
		l.Lock()
		if googleTokenService == nil {
			googleTokenService = &GoogleTokenService{}
		}
		l.Unlock()
	}
	return googleTokenService
}

//VerifyTotp 验证google验证码
func (googleTokenService GoogleTokenService) VerifyTotp(site string, accountName string, passcode string) bool {
	secureConfigureService := NewSecureConfigureService()
	googleSecret, err := secureConfigureService.Get(accountPrefix + accountName)
	if err != nil {
		Log.Error(err)
		return false
	}
	ok := totp.Validate(passcode, googleSecret)
	Log.WithField("account: ", accountPrefix+accountName).WithField("result:", ok).Debug("google token validate")
	return ok
}

//CreateGoogleToken 创建google令牌
func (googleTokenService GoogleTokenService) CreateGoogleToken(accountName string) (qr []byte, err error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "fs",
		AccountName: accountName,
	})
	if err != nil {
		Log.Error(err)
		return nil, err
	}

	secureConfigureService := NewSecureConfigureService()
	err = secureConfigureService.Set(accountPrefix+accountName, key.Secret())
	if err != nil {
		Log.Error(err)
		return nil, err
	}

	img, err := key.Image(200, 200)
	if err != nil {
		Log.Error(err)
		return nil, err
	}
	var buf bytes.Buffer
	png.Encode(&buf, img)
	return buf.Bytes(), nil
}
