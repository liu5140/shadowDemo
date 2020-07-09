package security

// TAnonymousAuthenticationToken Represents an anonymous Authentication.
var _ IAuthentication = &TAnonymousAuthenticationToken{}

type TAnonymousAuthenticationToken struct {
	Details       interface{}
	Authenticated bool
}

// NewAnonymousAuthenticationToken constructor for AnonymousAuthenticationToken
func NewAnonymousAuthenticationToken() *TAnonymousAuthenticationToken {
	return &TAnonymousAuthenticationToken{
		Authenticated: false,
	}
}

// GetPrincipal return the GetPrincipal being authenticated
func (anonymousAuthenticationToken *TAnonymousAuthenticationToken) GetPrincipal() string {
	return ANONYMOUS_ROLE
}

// GetCredential the credentials that prove the identity of the Principal
func (anonymousAuthenticationToken *TAnonymousAuthenticationToken) GetCredential() string {
	return ""
}

// GetDetails Records the details for the Authentication
func (anonymousAuthenticationToken *TAnonymousAuthenticationToken) GetDetails() interface{} {
	return anonymousAuthenticationToken.GetDetails
}

// SetDetails seter for Details
func (anonymousAuthenticationToken *TAnonymousAuthenticationToken) SetDetails(details interface{}) {
	// gob.Register(details)
	anonymousAuthenticationToken.Details = details
}

// IsAuthenticated getter for IsAuthenticated
func (anonymousAuthenticationToken *TAnonymousAuthenticationToken) IsAuthenticated() bool {
	return anonymousAuthenticationToken.Authenticated
}

// SetAuthenticated seter for Authenticated
func (anonymousAuthenticationToken *TAnonymousAuthenticationToken) SetAuthenticated(authenticationed bool) {
	anonymousAuthenticationToken.Authenticated = authenticationed
}
