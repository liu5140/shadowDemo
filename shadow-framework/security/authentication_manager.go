package security

type IAuthenticationManager interface {
	Authenticate(authentication IAuthentication) IAuthentication
}

type FAuthenticationManagerFactory func() IAuthenticationManager

var AuthenticationManagerFactories = make(map[string]FAuthenticationManagerFactory)

func RegisterAuthenticationManager(name string, factory FAuthenticationManagerFactory) {
	AuthenticationManagerFactories[name] = factory
}

func AuthenticationManagerInstance(name string) IAuthenticationManager {
	factory := AuthenticationManagerFactories[name]
	return factory()
}
