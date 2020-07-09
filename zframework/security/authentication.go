package security

// IAuthentication Represents the token for an authentication request or for an authenticated principal
type IAuthentication interface {
	// GetPrincipal return the GetPrincipal being authenticated
	GetPrincipal() string
	// Credentials the credentials that prove the identity of the Principal
	GetCredential() string
	GetDetails() interface{}
	SetDetails(interface{})
	IsAuthenticated() bool
}
