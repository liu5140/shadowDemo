package utils

import (
	"regexp"
	"strconv"
	"strings"
)

/*
   字符串保留小数位数，不会四舍五入，小数位数不足的补0
*/
func StringSaveBit(data string, n int) string {
	p := strings.LastIndex(data, ".")
	if p == -1 {
		return data + "." + strings.Repeat("0", n)
	}
	span := n - (len(data) - (p + 1))
	if span >= 0 {
		return data + strings.Repeat("0", span)
	} else {
		abs := 0 - span
		return data[0 : len(data)-abs]
	}
}

func String2int64(data string) (result int64) {
	if result, err := strconv.ParseInt(data, 10, 64); err == nil {
		return result
	}
	return
}

func String2float64(data string) (result float64) {
	if result, err := strconv.ParseFloat(data, 64); err == nil {
		return result
	}
	return
}



var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func LowerCaseFirstLetter(str string) string {
	if len(str) > 0 {
		first := strings.ToLower(string(str[0]))
		return first + string(str[1:])
	}
	return str
}