package service

import "testing"

// import (
// 	"testing"
// )

// func TestAliyunSendMessageCode(t *testing.T) {
// 	aliyunService := NewAliyunSmsService()
// 	err := aliyunService.SendMessageCode("86", "18503088056")
// 	if err != nil {
// 		t.Log(nil)
// 	}

// }

func TestSubStr(t *testing.T) {
	str := "OPRMW8CB312"
	// postscript := Substr(str, len(str)-5, len(str))
	postscript := str[len(str)-5:]

	t.Log(postscript)
}

//截取字符串 start 起点下标 end 终点下标(不包括)
func Substr(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		return ""
	}

	if end < 0 || end > length {
		return ""
	}
	return string(rs[start:end])
}
