package security

// IUserDetails Provides core user information.
type IUserDetails interface {
	GetUsername() string
	GetPassword() string
	IsAccountExpired() bool
	IsAccountLocked() bool
	IsCredentialsExpired() bool
}
