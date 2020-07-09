package security

var anonymousAuthenticationProvider *TAnonymousAuthenticationProvider

// TAnonymousAuthenticationProvider An AuthenticationProvider implementation that validates AnonymousAuthenticationTokens.
type TAnonymousAuthenticationProvider struct{}

func newAnonymousAuthenticationProvider() IAuthenticationProvider {
	if anonymousAuthenticationProvider == nil {
		anonymousAuthenticationProvider = new(TAnonymousAuthenticationProvider)
	}
	return anonymousAuthenticationProvider
}

func (provider *TAnonymousAuthenticationProvider) Authenticate(authentication IAuthentication) IAuthentication {
	if anonymousAuthenticationToken, ok := authentication.(*TAnonymousAuthenticationToken); ok {
		return anonymousAuthenticationToken
	}
	return nil
}
