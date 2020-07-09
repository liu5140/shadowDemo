package service

import (
	"bytes"
	"crypto/rand"
	"math/big"
	"shadowDemo/model"
	"shadowDemo/shadow-framework/idgenerator"
)

func GenAdminID() int64 {
	idGen := idgenerator.Instance()
	id := idGen.GenerateLongID(model.IDSpaceUser)
	return id
}

func GenAgentID() int64 {
	idGen := idgenerator.Instance()
	id := idGen.GenerateLongID(model.IDSpaceUser)
	return id
}

func GenPlayerID() int64 {
	idGen := idgenerator.Instance()
	id := idGen.GenerateLongID(model.IDSpaceUser)
	return id
}

func GenPersonalID() int64 {
	idGen := idgenerator.Instance()
	id := idGen.GenerateLongID(model.IDSpaceUser)
	return id
}

func GenOrderID() int64 {
	idGen := idgenerator.Instance()
	id := idGen.GenerateLongID(model.IDSpaceOrder)
	return id
}

//推广码ID
func GenRecommendCode() string {
	idGen := idgenerator.Instance()
	code := idGen.GenerateGuardID(model.IDSpaceRecommendCode)
	return code
}

func GetRandomString(len int) string {
	var container string
	var str = "ABCDEFGHIJKLMNPQRSTUVWXYZ123456789"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}
