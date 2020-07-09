package security

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	encoder := &TDefaultPasswordEncoder{}
	encodedString := encoder.Encode("secret")
	t.Log(encodedString)
	assert.NotEmpty(t, encodedString)

}

func TestMatches(t *testing.T) {
	encoder := &TDefaultPasswordEncoder{}
	hash := encoder.Encode("secret")
	assert.True(t, encoder.Matches("secret", hash))
}
