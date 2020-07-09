package middleware

type IUrlRegistry interface {
	GetPath() string
}

//FUrlRegistryFactory UrlRegistry factory
type FUrlRegistryFactory func() IUrlRegistry

var UrlRegistryFactories = make(map[string]FUrlRegistryFactory)

//RegisterUrlRegistry RegisterUrlRegistry
func RegisterUrlRegistry(name string, factory FUrlRegistryFactory) {
	UrlRegistryFactories[name] = factory
}

//UrlRegistryInstance UrlRegistryInstance
func UrlRegistryInstance(name string) IUrlRegistry {
	factory := UrlRegistryFactories[name]
	return factory()
}

type TDefaultLoginUrlRegistry struct{}

func newDefaultLoginUrlRegistry() IUrlRegistry {
	return &TDefaultLoginUrlRegistry{}
}

func (loginUrlRegistry *TDefaultLoginUrlRegistry) GetPath() string {
	return "/login"
}

type TDefaultLogoutUrlRegistry struct{}

func newDefaultLogoutUrlRegistry() IUrlRegistry {
	return &TDefaultLogoutUrlRegistry{}
}

func (logoutUrlRegistry *TDefaultLogoutUrlRegistry) GetPath() string {
	return "/logout"
}
