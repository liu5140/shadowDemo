package security

// TRequestAuthenticationToken Represents an casbin requets Authentication.
type TRequestAuthenticationToken struct {
	Details       interface{}
	Authenticated bool
}

var _ IAuthentication = &TRequestAuthenticationToken{}

// NewRequestAuthenticationToken constructor for RequestAuthenticationToken
func NewRequestAuthenticationToken() *TRequestAuthenticationToken {
	return &TRequestAuthenticationToken{
		Authenticated: false,
	}
}

// GetPrincipal return the GetPrincipal being authenticated
func (RequestAuthenticationToken *TRequestAuthenticationToken) GetPrincipal() string {
	return ""
}

// GetCredential the credentials that prove the identity of the Principal
func (RequestAuthenticationToken *TRequestAuthenticationToken) GetCredential() string {
	return ""
}

// GetDetails Records the details for the Authentication
func (RequestAuthenticationToken *TRequestAuthenticationToken) GetDetails() interface{} {
	return RequestAuthenticationToken.Details
}

// SetDetails seter for Details
func (RequestAuthenticationToken *TRequestAuthenticationToken) SetDetails(details interface{}) {
	RequestAuthenticationToken.Details = details
}

// IsAuthenticated getter for IsAuthenticated
func (RequestAuthenticationToken *TRequestAuthenticationToken) IsAuthenticated() bool {
	return RequestAuthenticationToken.Authenticated
}

// SetAuthenticated seter for Authenticated
func (RequestAuthenticationToken *TRequestAuthenticationToken) SetAuthenticated(authenticationed bool) {
	RequestAuthenticationToken.Authenticated = authenticationed
}
