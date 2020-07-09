package middleware

import (
	"net/http"

	"shadowDemo/shadow-framework/security"

	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/i18n"
)

type (
	WrongGoogleToken struct {
		Err error
	}
	WrongCaptchaCode struct {
		Err error
	}
	MissGoogleToken struct {
		Err error
	}
	VisitLimit struct {
		Err error
	}
	RoleAmountLimit struct {
		Err error
	}

	BalanceFrozenErr struct {
		Err error
	}

	AccountAuditingErr struct {
		Err error
	}

	JwtErr struct {
		Err error
	}
)

func (e WrongGoogleToken) Error() string {
	return e.Err.Error()
}
func (e WrongCaptchaCode) Error() string {
	return e.Err.Error()
}
func (e MissGoogleToken) Error() string {
	return e.Err.Error()
}
func (e VisitLimit) Error() string {
	return e.Err.Error()
}
func (e RoleAmountLimit) Error() string {
	return e.Err.Error()
}

func (e BalanceFrozenErr) Error() string {
	return e.Err.Error()
}
func (e AccountAuditingErr) Error() string {
	return e.Err.Error()
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.Errors) > 0 {
			T := c.MustGet("T").(i18n.TranslateFunc)
			for _, err := range c.Errors {
				Log.Infoln("===========", err)

				Log.Warn(err)
				switch err.Err.(type) {
				case security.WrongUserNamePasswordError:
					newGenError(c, T("key_alert_error_username_password_error"))
				case security.NotPromissionError:
					newGenError(c, T("key_alert_error_not_promission_error"))
				case security.AccountLockedError:
					newGenError(c, T("key_alert_error_account_locked_error"))
				case security.AccountExpiredError:
					newGenError(c, T("key_alert_error_account_expired_error"))
				case security.SmsExpiredError:
					newGenError(c, T("key_alert_error_sms_expired_error"))
				case WrongGoogleToken:
					newGenError(c, T("key_alert_error_token_error"))
				case WrongCaptchaCode:
					newGenError(c, T("key_alert_error_captcha_error"))
				case MissGoogleToken:
					newGenError(c, T("key_alert_error_captcha_error"))
				case VisitLimit:
					newGenError(c, T("key_alert_ip_not_allowed_error"))
				case RoleAmountLimit:
					newGenError(c, T("key_alert_error_amount_not_promission_error"))
				case security.JwtExpiredErr:
					newForbiddenError(c, "TOKEN 失效")
				}
				break
			}
		}
	}
}

// swagger:model
type GenericMessageBody struct {
	Msg  string
	Code int
}

func newGenError(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, GenericMessageBody{
		Msg:  message,
		Code: http.StatusBadRequest,
	})
	c.Abort()
}

func newForbiddenError(c *gin.Context, message string) {
	c.JSON(http.StatusForbidden, GenericMessageBody{
		Msg:  message,
		Code: http.StatusUnauthorized,
	})
	c.Abort()
}
