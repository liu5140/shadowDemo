package utils

import "encoding/base64"

const (
	base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)

var coder = base64.NewEncoding(base64Table)

func Base64Encode(src []byte) []byte {
	return []byte(coder.EncodeToString(src))
}

func Base64Decode(src []byte) ([]byte, error) {
	return coder.DecodeString(string(src))
}

// base64 encode
func Base64Encode2(str []byte) string {
	return base64.StdEncoding.EncodeToString(str)
}

// base64 Std decode
func Base64Decode2(str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(str)
}
