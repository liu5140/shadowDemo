package security

// TRbacDomainCasbinAuthenticationRequestResolver default implementetion for rbac domain based casbin authentication request
type TRbacDomainCasbinAuthenticationRequestResolver struct{}

// newRbacDomainCasbinAuthenticationRequestResolver Constructor method for CasbinAuthenticationRequestResolver
func newRbacDomainCasbinAuthenticationRequestResolver() CasbinAuthenticationRequestResolver {
	return &TRbacDomainCasbinAuthenticationRequestResolver{}
}

func (rbacDomainCasbinAuthenticationRequestResolver TRbacDomainCasbinAuthenticationRequestResolver) ObtainCasbinRequest(domain string, requestURI string, method string, principal string) TCasbinPolicyDetails {
	return TCasbinPolicyDetails{
		Domain:  domain,
		Sub:     principal,
		Obj:     requestURI,
		Act:     method,
		Service: "",
		Eft:     "",
	}
}
