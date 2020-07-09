package i18nmiddleware

import (
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/i18n"
	"github.com/sirupsen/logrus"
)

//I18nResolver set locale and lang parameters for the request context
func I18nResolver() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookieLang, _ := c.Cookie("lang")
		acceptLang := c.GetHeader("Accept-Language")
		xlang := c.GetHeader("applang")
		Log.Infoln("====Lang=========", acceptLang)
		Log.Infoln("====xlang=========", xlang)
		defaultLang := "zh"
		T, lang := i18n.MustTfuncAndLanguage(cookieLang, acceptLang, defaultLang)
		Log.WithFields(logrus.Fields{
			"cookieLang":  cookieLang,
			"acceptLang":  acceptLang,
			"defaultLang": defaultLang,
			"usingLang":   lang,
		}).Debug("I18nResolver")
		c.Set("T", T)
		c.Set("Lang", lang)
		c.Next()
	}
}
