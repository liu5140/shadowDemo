package utils

import (
	"encoding/hex"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"time"

	. "shadowDemo/zframework/logger"
)

const (
	TIME_FORMAT_WITH_MS         = "2006-01-02 15:04:05.000"
	TIME_FORMAT                 = "2006-01-02 15:04:05"
	TIME_FORMAT_COMPACT         = "20060102150405"
	TIME_FORMAT_WITH_MS_COMPACT = "20060102150405.000"
	DATE_FORMAT                 = "2006-01-02"
	DATE_FORMAT_COMPACT         = "20060102"
	MONTH_FORMAT                = "2006-01"
)

func ASSERT(exp bool, info ...string) { // 接受一个字符串参数
	if !exp {
		infostr := ""
		if len(info) > 0 {
			infostr = info[0]
		}
		Log.Errorf("ASSERT FAILED!\ninfo=[%v]\nstack = [%v]\n", infostr, string(debug.Stack()))
		panic("ASSERT FAILED")
	}
}

func CatchPanic() {
	if err := recover(); err != nil {
		Log.Errorf("panic !!! err = %v ", err)
	}
}

func CatchPanicWarning() {
	if err := recover(); err != nil {
		Log.Warnf("panic !!! err = %v ", err)
	}
}

func CatchException() {
	if err := recover(); err != nil {
		fullPath, _ := exec.LookPath(os.Args[0])
		fname := filepath.Base(fullPath)

		datestr := NowDateStr()
		outstr := fmt.Sprintf("\n======\n[%v] err=%v, stack=%v\n======\n", time.Now(), err, string(debug.Stack()))
		filename := "./log/panic_" + fname + datestr + ".log"
		f, err2 := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
		ASSERT(err2 == nil)
		defer f.Close()
		f.WriteString(outstr)

		Log.Errorf("err = %v ", err)
	}
}

func NowDateStr() string {
	timenow := time.Now().Format(DATE_FORMAT)
	return timenow
}

func UseMaxCpu() {
	// multiple cups using
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func HexBuffer(buffer []byte) string {
	s := hex.EncodeToString(buffer)
	n := 8
	m := 8
	c := 0
	slen := len(s)
	if slen%n == 0 {
		c = slen / n
	} else {
		c = (slen / n) + 1
	}

	res := ""
	for i := 0; i < c; i++ {
		res = res + s[i:i+8] + " "
		if (i+1)%m == 0 {
			res = res + "\n"
		}
	}

	res = fmt.Sprintf("\n=======================================================================\n%v\n=======================================================================\n", res)
	return res
}

func GetProgName() string {
	fullPath, _ := exec.LookPath(os.Args[0])
	fname := filepath.Base(fullPath)

	return fname
}

func EncodeURI(data string) string {
	return url.QueryEscape(data)
}

func DecodeURI(data string) (string, error) {
	sdata, err := url.QueryUnescape(data)
	if err != nil {
		Log.Errorf("url.QueryUnescape err = %v", err)
		return "", err
	}
	return sdata, nil
}

func IsNullTime(t time.Time) bool {
	year := t.Year()
	if year == 1 {
		return true
	} else {
		return false
	}
}

// IsLower check letter is lower case or not
func IsLower(b byte) bool {
	return b >= 'a' && b <= 'z'
}

// IsUpper check letter is upper case or not
func IsUpper(b byte) bool {
	return b >= 'A' && b <= 'Z'
}

// IsLetter check character is a letter or not
func IsLetter(b byte) bool {
	return IsLower(b) || IsUpper(b)
}

// IsNumber check character is a number or not
func IsNumber(c byte) bool {
	return c >= '0' && c <= '9'
}

// IsAlNum check character is a alnum or not
func IsAlNum(c byte) bool {
	return IsLetter(c) || IsLetter(c)
}

func TimeStr(t time.Time) string {
	return t.Format(TIME_FORMAT)
}
