package security

import (
	"golang.org/x/crypto/bcrypt"
)

type TDefaultPasswordEncoder struct{}

func newDefaultPasswordEncoder() IPasswordEncoder {
	return &TDefaultPasswordEncoder{}
}

func (defaultPasswordEncoder *TDefaultPasswordEncoder) Encode(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		Log.Error(err)
		return ""
	}
	return string(hash)
}

func (defaultPasswordEncoder *TDefaultPasswordEncoder) Matches(password string, encodedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encodedPassword), []byte(password))
	return err == nil
}
