package security

// TUsernamePasswordAuthenticationToken An Authentication implementation that is designed for simple presentation of a username and password.
type TUsernamePasswordAuthenticationToken struct {
	Pprincipal     string
	Pcredential    string
	Pdetails       interface{}
	Pauthenticated bool
}

var _ IAuthentication = &TUsernamePasswordAuthenticationToken{}

// NewUsernamePasswordAuthenticationToken constructor for UsernamePasswordAuthenticationToken
func NewUsernamePasswordAuthenticationToken(principal string, credential string) *TUsernamePasswordAuthenticationToken {
	return &TUsernamePasswordAuthenticationToken{
		Pprincipal:     principal,
		Pcredential:    credential,
		Pauthenticated: false,
	}
}

// GetPrincipal return the GetPrincipal being authenticated
func (usernamePasswordAuthenticationToken *TUsernamePasswordAuthenticationToken) GetPrincipal() string {
	return usernamePasswordAuthenticationToken.Pprincipal
}

// GetCredential the credentials that prove the identity of the Principal
func (usernamePasswordAuthenticationToken *TUsernamePasswordAuthenticationToken) GetCredential() string {
	return usernamePasswordAuthenticationToken.Pcredential
}

// GetDetails Records the details for the Authentication
func (usernamePasswordAuthenticationToken *TUsernamePasswordAuthenticationToken) GetDetails() interface{} {
	return usernamePasswordAuthenticationToken.Pdetails
}

// SetDetails seter for Details
func (usernamePasswordAuthenticationToken *TUsernamePasswordAuthenticationToken) SetDetails(details interface{}) {
	usernamePasswordAuthenticationToken.Pdetails = details
}

// IsAuthenticated getter for IsAuthenticated
func (usernamePasswordAuthenticationToken *TUsernamePasswordAuthenticationToken) IsAuthenticated() bool {
	return usernamePasswordAuthenticationToken.Pauthenticated
}

// SetAuthenticated seter for Authenticated
func (usernamePasswordAuthenticationToken *TUsernamePasswordAuthenticationToken) SetAuthenticated(authenticationed bool) {
	usernamePasswordAuthenticationToken.Pauthenticated = authenticationed
}
