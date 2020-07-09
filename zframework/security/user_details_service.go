package security

// IUserDetailService the interface requires only one read-only method, which simplifies support for new data-access strategies.
type IUserDetailService interface {
	LoadUserByUsername(username string) IUserDetails
}

type FUserDetailServiceFactory func() interface{}

var userDetailServiceFactories = make(map[string]FUserDetailServiceFactory)

func RegisterUserDetailService(name string, factory FUserDetailServiceFactory) {
	userDetailServiceFactories[name] = factory
}

func CreateUserDetailService(name string) IUserDetailService {
	factory := userDetailServiceFactories[name]
	return factory().(IUserDetailService)
}
