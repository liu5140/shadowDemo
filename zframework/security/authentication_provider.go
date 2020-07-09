package security

// IAuthenticationProvider Indicates a struct can process a specific
type IAuthenticationProvider interface {
	Authenticate(authentication IAuthentication) IAuthentication
}

type FAuthenticationProviderFactory func() IAuthenticationProvider

var authenticationProviderFactories = make(map[string]FAuthenticationProviderFactory)

func RegisterAuthenticationProvider(name string, factory FAuthenticationProviderFactory) {
	authenticationProviderFactories[name] = factory
}

func AuthenticationProviderInstance(name string) IAuthenticationProvider {
	factory := authenticationProviderFactories[name]
	return factory()
}

func AuthenticationProviderInstances() []IAuthenticationProvider {
	providers := make([]IAuthenticationProvider, len(authenticationProviderFactories))
	var i int
	for _, factory := range authenticationProviderFactories {
		providers[i] = factory()
		i++
	}
	return providers
}
