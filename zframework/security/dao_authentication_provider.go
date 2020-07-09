package security

import (
	"github.com/sirupsen/logrus"
)

var daoAuthenticationProvider *TDaoAuthenticationProvider

// TDaoAuthenticationProvider an implementation that retrieves user details from a UserDetailService
type TDaoAuthenticationProvider struct {
	userDetailService IUserDetailService
	passwordEncoder   IPasswordEncoder
}

func newDaoAuthenticationProvider() IAuthenticationProvider {
	if daoAuthenticationProvider == nil {
		daoAuthenticationProvider = &TDaoAuthenticationProvider{
			userDetailService: CreateUserDetailService(USER_DETAILS_SERVICE),
			passwordEncoder:   PasswordEncoderInstance(PASSWORD_ENCODER),
		}
	}
	return daoAuthenticationProvider
}

// SetPasswordEncoder set password encode
func (provider *TDaoAuthenticationProvider) SetPasswordEncoder(encoder IPasswordEncoder) {
	provider.passwordEncoder = encoder
}

func (provider *TDaoAuthenticationProvider) SetUserDetailService(userDetailService IUserDetailService) {
	provider.userDetailService = userDetailService
}

func (provider *TDaoAuthenticationProvider) Authenticate(authentication IAuthentication) IAuthentication {
	if usernamePasswordAuthenticationToken, ok := authentication.(*TUsernamePasswordAuthenticationToken); ok {
		userDetails := provider.userDetailService.LoadUserByUsername(authentication.GetPrincipal())
		usernamePasswordAuthenticationToken.SetDetails(userDetails)
		if userDetails != nil && !userDetails.IsAccountExpired() && !userDetails.IsAccountLocked() && !userDetails.IsCredentialsExpired() {
			if provider.authenticationChecks(userDetails, usernamePasswordAuthenticationToken) {
				usernamePasswordAuthenticationToken.SetAuthenticated(true)
			}
		}
		return usernamePasswordAuthenticationToken
	}
	return nil
}

func (provider *TDaoAuthenticationProvider) authenticationChecks(userDetails IUserDetails, authentication *TUsernamePasswordAuthenticationToken) bool {
	if authentication.GetCredential() == "" {
		return false
	}
	if provider.passwordEncoder.Matches(authentication.GetCredential(), userDetails.GetPassword()) {
		Log.WithFields(logrus.Fields{
			"username": authentication.GetPrincipal(),
			"IsMatch":  true,
		}).Debug("password check debug")
		return true
	}
	Log.WithFields(logrus.Fields{
		"username": authentication.GetPrincipal(),
		"IsMatch":  false,
	}).Debug("password check debug")
	return false
}
