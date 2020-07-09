package security

// TUser default implements for Userdetails interface
type TUser struct {
	Username           string `form:"username" json:"username" xml:"username"  binding:"required"`
	Password           string `form:"password" json:"password" xml:"password"  binding:"required"`
	AccountExpired     bool
	AccountLocked      bool
	CredentialsExpired bool
}

// GetUsername getter for username
func (user *TUser) GetUsername() string {
	return user.Username
}

// SetUsername setter for username
func (user *TUser) SetUsername(username string) {
	user.Username = username
}

// GetPassword getter for password
func (user *TUser) GetPassword() string {
	return user.Password
}

// SetPassword setter for password
func (user *TUser) SetPassword(password string) {
	user.Password = password
}

// IsAccountExpired getter for AccountExpired
func (user *TUser) IsAccountExpired() bool {
	return user.AccountExpired
}

// SetAccountExpired setter for AccountExpired
func (user *TUser) SetAccountExpired(accountExpired bool) {
	user.AccountExpired = accountExpired
}

// IsAccountLocked Getter for AccountLocked
func (user *TUser) IsAccountLocked() bool {
	return user.AccountLocked
}

// SetAccountLocked setter for AccountLocked
func (user *TUser) SetAccountLocked(accountLocked bool) {
	user.AccountLocked = accountLocked
}

// IsCredentialsExpired getter for CredentialsExpired
func (user *TUser) IsCredentialsExpired() bool {
	return user.CredentialsExpired
}

// SetCredentialsExpired setter for CredentialsExpired
func (user *TUser) SetCredentialsExpired(credentialsExpired bool) {
	user.CredentialsExpired = credentialsExpired
}
