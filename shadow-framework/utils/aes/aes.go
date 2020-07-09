package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"fmt"
	"net/url"
	"shadowDemo/shadow-framework/utils"
)

// /*
// 	算法: AES-128, AES-192, or AES-256
// 	模式: CBC
// 	填充: PKCS5
// 	偏移量: key值前16位
// 	输出: hex
// */
func AESEncryptStr(origDataStr string, key string) (string, error) {
	eCode, err := AesEncrypt([]byte(origDataStr), []byte(key))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(eCode), nil
}

func AESDecryptStr(cryptedDataStr string, key string) (string, error) {
	retByte, err := hex.DecodeString(cryptedDataStr)
	if err != nil {
		return "", err
	}
	dCode, err := AesDecrypt(retByte, []byte(key))
	if err != nil {
		return "", err
	}
	return string(dCode), err
}

func AESEncryptStr2(origDataStr string, key string) (string, error) {
	eCode, err := AesEncrypt([]byte(origDataStr), []byte(key))
	if err != nil {
		return "", err
	}

	b64str := url.QueryEscape(string(utils.Base64Encode(eCode)))

	return b64str, nil
}

func AESDecryptStr2(cryptedDataStr string, key string) (string, error) {

	unesc, err := url.QueryUnescape(cryptedDataStr)
	if err != nil {
		return "", err
	}

	retByte, err := utils.Base64Decode([]byte(unesc))
	if err != nil {
		return "", err
	}

	dCode, err := AesDecrypt(retByte, []byte(key))
	if err != nil {
		return "", err
	}
	return string(dCode), err
}

func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData, err = PKCS5UnPadding(origData)
	if err != nil {
		return nil, err
	}
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
func PKCS5UnPadding(origData []byte) ([]byte, error) {
	length := len(origData)
	unpadding := int(origData[length-1])
	if length <= unpadding {
		return nil, errors.New("unpadding error!")
	}
	return origData[:(length - unpadding)], nil
}

const (
	InternalCryptKey = "Internaloooo0000JjKkIiBbAaXxCcVv"
)

func EncryptHex(origDataStr string) string {
	s, err := InternalEncryptStr(origDataStr)
	if err != nil {
		return ""
	}
	return s
}

func DecryptStr(origDataStr string) string {
	s, err := InternalDecryptStr(origDataStr)
	if err != nil {
		return ""
	}
	return s
}

func InternalEncryptStr(origDataStr string) (string, error) {
	eCode, err := AesEncrypt([]byte(origDataStr), []byte(InternalCryptKey))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(eCode), nil
}

func InternalDecryptStr(cryptedDataStr string) (string, error) {
	retByte, _ := hex.DecodeString(cryptedDataStr)
	dCode, err := AesDecrypt(retByte, []byte(InternalCryptKey))
	if err != nil {
		return "", err
	}
	return string(dCode), err
}

// AES ECB Mode
func AesEcbEncrypt(origData, key []byte) ([]byte, error) {
	fmt.Println("========", key)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	fmt.Println("========", block)

	// ECB Mode
	bs := block.BlockSize()
	origData = PKCS5Padding(origData, bs)
	if len(origData)%bs != 0 {
		return nil, errors.New("Need a multiple of the blocksize")
	}
	out := make([]byte, len(origData))
	dst := out
	for len(origData) > 0 {
		block.Encrypt(dst, origData[:bs])
		origData = origData[bs:]
		dst = dst[bs:]
	}
	return out, nil
}

func AesEcbDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// ECB Mode
	out := make([]byte, len(crypted))
	dst := out
	bs := block.BlockSize()
	if len(crypted)%bs != 0 {
		return nil, errors.New("crypto/cipher: input not full blocks")
	}
	for len(crypted) > 0 {
		block.Decrypt(dst, crypted[:bs])
		crypted = crypted[bs:]
		dst = dst[bs:]
	}
	// out = ZeroUnPadding(out)
	out, err = PKCS5UnPadding(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

/*
	算法: AES-128, AES-192, or AES-256
	模式: CBC
	填充: 数据已填充
	偏移量: iv 指定
	输出: base64
*/
func AesEncryptToBase64(padData, key, iv []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(padData))
	blockMode.CryptBlocks(crypted, padData)

	return utils.Base64Encode2(crypted), nil
}

/*
	算法: AES-128, AES-192, or AES-256
	模式: ECB
	填充: PKCS5Padding
	偏移量:
	输出: base64
*/
func AesEcbEncryptToBase64(origData, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// ECB Mode
	bs := block.BlockSize()
	origData = PKCS5Padding(origData, bs)
	if len(origData)%bs != 0 {
		return "", errors.New("Need a multiple of the blocksize")
	}
	out := make([]byte, len(origData))
	dst := out
	for len(origData) > 0 {
		block.Encrypt(dst, origData[:bs])
		origData = origData[bs:]
		dst = dst[bs:]
	}

	return utils.Base64Encode2(out), nil
}
