package utils

import (
	"math/rand"
	"sync"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

var l sync.Mutex

func RandStringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	l.Lock()
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	l.Unlock()
	return string(b)
}

func RandString(length int) string {
	return RandStringWithCharset(length, charset)
}
