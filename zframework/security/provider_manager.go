package security

var providerManager IAuthenticationManager

type TProviderManager struct {
	providers []IAuthenticationProvider
}

// InitAuthenticationManager initial a InitAuthenticationManager instance
func newAuthenticationManager() IAuthenticationManager {
	if providerManager == nil {
		providerManager = &TProviderManager{
			providers: AuthenticationProviderInstances(),
		}
	}
	return providerManager
}

// Authenticate Attempts to authenticate the passed Authentication object, returning a fully populated Authentication object (including granted authorities) if successful.
func (providerManager *TProviderManager) Authenticate(authentication IAuthentication) IAuthentication {
	var result IAuthentication
	for _, provider := range providerManager.providers {
		result = provider.Authenticate(authentication)
		if result != nil {
			authentication.SetDetails(result.GetDetails())
			break
		}
	}
	return result
}

func (providerManager *TProviderManager) AuthenticationProvider() []IAuthenticationProvider {
	return providerManager.providers
}

func (providerManager *TProviderManager) SetAuthenticationProvider(providers []IAuthenticationProvider) {
	providerManager.providers = providers
}
