package middleware

import (
	"bytes"
	"errors"
	"image"
	"image/png"
	"net/http"
	"os"
	"shadowDemo/zframework/utils/encoder"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/hanguofeng/gocaptcha"
)

var CaptchaCtrl TCaptchaCtrl

type CaptchaRes struct {
	Image string
	Key   string
}

func CaptchaHandler(getCaptchaURL string, context string) gin.HandlerFunc {
	CaptchaInit()
	var secureURL = map[string]string{
		context + "/login":  http.MethodPost,
		context + "/player": http.MethodPost,
	}
	return func(c *gin.Context) {
		if c.Request.URL.RequestURI() == getCaptchaURL && c.Request.Method == http.MethodPost {
			key, img, err := CaptchaGetKeyAndImage()
			if err != nil {
				Log.Error(err)
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			c.JSON(http.StatusOK, CaptchaRes{
				Image: img,
				Key:   key,
			})
			c.Abort()
			return
		}

		if v, ok := secureURL[c.Request.URL.Path]; ok && c.Request.Method == v {
			code := c.GetHeader("CaptchaCode")
			key := c.GetHeader("CaptchaKey")

			if code == "" || key == "" {
				c.Error(WrongCaptchaCode{
					Err: errors.New("ivalid captcha code"),
				})
				return
			}

			_, err := CaptchaVerifyCode(key, code)
			if err != nil {
				c.Error(WrongCaptchaCode{
					Err: err,
				})
			}
		}

	}

}

// code 之外的验证信息, 用于加强验证码验证
type TExtraVerifyItem struct {
	UserName  string
	ForMethod string
}

type TCaptchaCtrl struct {
	Captcha        *gocaptcha.Captcha
	ExtraVerifyMap map[string]TExtraVerifyItem
	Mutex          sync.Mutex
	Disabled       bool
}

func CaptchaInit() {
	captcha, err := gocaptcha.CreateCaptchaFromConfigFile("./config/gocaptcha/gocaptcha.conf")
	if nil != err {
		Log.Error("config load failed:%s", err.Error())
		os.Exit(-7)
	} else {
		CaptchaCtrl.Captcha = captcha
		CaptchaCtrl.ExtraVerifyMap = make(map[string]TExtraVerifyItem)
		CaptchaCtrl.Disabled = false
	}
}

func CaptchaGetKeyAndImage() (string, string, error) {
	CaptchaCtrl.Mutex.Lock()
	defer CaptchaCtrl.Mutex.Unlock()

	key, err := CaptchaCtrl.Captcha.GetKey(4)
	if err != nil {
		return "", "", err
	}

	image, err := CaptchaCtrl.Captcha.GetImage(key)
	if err != nil {
		return "", "", err
	}

	buff := new(bytes.Buffer)
	err = png.Encode(buff, image)
	if err != nil {
		return "", "", err
	}

	imgb64 := encoder.Base64Encode(buff.Bytes())
	imgstr := "data:image/png;base64," + string(imgb64)

	return key, imgstr, nil
}

func CaptchaGetImage(key string) (image.Image, error) {
	return CaptchaCtrl.Captcha.GetImage(key)
}

func CaptchaVerifyCode(key string, code string) (bool, error) {

	if CaptchaCtrl.Disabled {
		return true, nil
	}

	CaptchaCtrl.Mutex.Lock()
	defer CaptchaCtrl.Mutex.Unlock()
	// defer delete(CaptchaCtrl.ExtraVerifyMap, key) // 不管是否验证成功,清理验证码,保证其只能用一次

	// _, exists := CaptchaCtrl.ExtraVerifyMap[key]
	// if !exists {
	// 	return false, errors.New("Invalid request for verification")
	// }

	ok, msg := CaptchaCtrl.Captcha.Verify(key, code)
	if !ok {
		return false, errors.New(msg)
	}

	return true, nil
}
