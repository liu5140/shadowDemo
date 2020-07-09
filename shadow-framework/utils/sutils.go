package utils

import (
	"bytes"
	"crypto/md5"
	"crypto/rc4"
	"encoding/hex"
	"fmt"
	"net"
	"os"
	"os/signal"
	"regexp"
	"runtime/pprof"
	. "shadowDemo/shadow-framework/logger"
	"strconv"
	"strings"
	"syscall"
	"unicode"
)

type StringBuffer bytes.Buffer

type outputInterface interface {
	PutStruct(debug *StringBuffer) *StringBuffer
}

func MD5(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (this *StringBuffer) String() string {
	thisBuffer := (*bytes.Buffer)(this)
	return thisBuffer.String()
}

func GetMidStr(str string, s string, e string) (midstr string) {
	si := strings.Index(str, s)
	if si < 0 {
		return ""
	}

	sp := si + len(s)
	str2 := str[sp:]
	ei := strings.Index(str2, e)

	if ei < 0 {
		return str[sp:]
	} else {
		return str[sp : sp+ei]
	}
}

func Substr(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		return "" //panic("start is wrong")
	}

	if end < 0 || end > length {
		return "" //panic("end is wrong")
	}

	return string(rs[start:end])
}

func SignalHandler() {

	Log.Info("SignalHandler starting... ...")

	//defer CatchException()

	c := make(chan os.Signal)

	Log.Debug("signal notify")
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	Log.Debug("Recieved sig from channel")
	for sig := range c {
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP:
			Log.Warn(" Signal = %v", sig)
			//                if UsePprof == "1" {
			//                    Log.Info("Will stop Pprof ... ...")
			//                    StopPprof()
			//                }
		}
	}
}

func IsValidIp4(ipaddr string) bool {
	testInput := net.ParseIP(ipaddr)
	if testInput.To4() == nil {
		return false
	}

	return true
}

func Str2Hex(bstr string) []byte {
	length := len(bstr)
	cnt := length / 2
	var bout []byte = make([]byte, cnt)
	j := 0
	for i := 0; i < length; i = i + 2 {
		s := string(bstr[i : i+2])
		//    		Log.Debug("s = %v", s)
		fmt.Sscanf(s, "%02x", &bout[j])
		//    		Log.Debug("bout[%v] = %02x", j, bout[j])
		j++
	}

	return bout
}

func SimpleXor(key string, src string) string {

	bsrc := Str2Hex(src)

	if len(src) > (2 * len(key)) {
		return ""
	}

	var dst []byte

	for i := 0; i < len(bsrc); i++ {
		d := bsrc[i] ^ key[i]
		//dst = dst + d
		dst = append(dst, d)
	}

	return string(dst)
}

func SimpleXorStr(key string, src string) string {
	if len(src) > len(key) {
		return ""
	}

	var dst []byte

	for i := 0; i < len(src); i++ {
		d := src[i] ^ key[i]
		//dst = dst + d
		dst = append(dst, d)
	}

	dstr := ""
	for i := 0; i < len(dst); i++ {
		ds := fmt.Sprintf("%02x", dst[i])
		dstr = dstr + ds
	}

	return dstr
}

func RC4Encrypt(key []byte, data []byte) ([]byte, error) {
	c, err := rc4.NewCipher(key)
	if err != nil {
		Log.Error("rc4 init key failed, err = %v", err)
		return nil, err
	}

	encrypted := make([]byte, len(data))
	c.XORKeyStream(encrypted, data)
	c.Reset()

	return encrypted, nil
}

func RC4Decrypt(key []byte, enc_data []byte) ([]byte, error) {
	c, err := rc4.NewCipher(key)
	if err != nil {
		Log.Error("rc4 init key failed, err = %v", err)
		return nil, err
	}
	decrypted := make([]byte, len(enc_data))
	c.XORKeyStream(decrypted, enc_data)
	c.Reset()

	return decrypted, nil
}

func FastEncrypt(key string, data string) string {

	enckey, err := RC4Encrypt([]byte(key), []byte(data))
	if err != nil {
		Log.Error("rc4 encrypt failed, key = %v, data = %v", key, data)
		return ""
	}

	encstr := fmt.Sprintf("%x", enckey)

	return encstr
}

func FastDecrypt(key string, encData string) string {

	dc, err := hex.DecodeString(encData)
	if err != nil {
		Log.Error("hex decoding failed, encData = %v, err = %v", encData, err)
		return ""
	}

	dec, err := RC4Decrypt([]byte(key), []byte(dc))
	if err != nil {
		Log.Error("rc4 decrypt failed, key = %v, encData = %v", key, encData)
		return ""
	}

	return string(dec)
}

var proffile *os.File

// func StartPprof() {
// 	filename := os.Args[0] + NowStr() + ".prof"
// 	proffile, err := os.Create(filename)
// 	if err != nil {
// 		Log.Error("create prof file failed!, err=%v", err)
// 		return
// 	}

// 	pprof.StartCPUProfile(proffile)
// }

func StopPprof() {
	pprof.StopCPUProfile()
	proffile.Close()
}

////////////////////////////////////////
////////////////////////////////////////
////////////////////////////////////////
///// for param check util functions

func IsDigital(s string) bool {
	_, err := strconv.Atoi(s)
	if err != nil {
		return false
	}

	return true
}

func IsValidInput(s string) bool {
	re := `[a-z0-9A-Z_]+\-?[a-z0-9A-Z_]*`
	reg := regexp.MustCompile(re)
	ss := reg.FindAllString(s, -1)
	if len(ss) == 1 {
		return true
	} else {
		return false
	}
}

func HasRepeat(s string) bool {
	for _, c := range s {
		n := strings.Count(s, string(c))
		if n > 1 {
			return true
		}
	}
	return false
}

func Trim(s string) string {
	return strings.Trim(s, " \n\r\t")
}

func MaxRepeatCount(s string) int {

	maxCount := 0

	for _, c := range s {
		n := strings.Count(s, string(c))
		if n > maxCount {
			maxCount = n
		}
	}

	return maxCount
}

func HasSame(n0 []string, n1 []string) bool {
	for _, n1x := range n1 {
		for _, n0x := range n0 {
			if n1x == n0x {
				return true
			}
		}
	}
	return false
}

func HasSameChar(n0 string, n1 string) bool {
	for _, n1x := range n1 {
		for _, n0x := range n0 {
			if n1x == n0x {
				return true
			}
		}
	}

	return false
}

func StringSplit(s string, sep string) []string {
	if len(s) == 0 {
		return []string{}
	}
	ss := strings.Split(s, sep)
	return ss
}

////////////////////////////////////////
//////////字符串检查函数////////////////
///////////////////////////////////////

func IsAlphabetNumber(s string) bool {
	reg := regexp.MustCompile(`[\W]+`)
	return !reg.MatchString(s)
}

// ' `(x60) ^ <script>
func IsLegalMsgString(s string) bool {
	reg := regexp.MustCompile(`['\x60\^]+|(?i:\<script\>)`)
	return !reg.MatchString(s)
}

func IsLegalChineseUserName(s string) bool {
	for _, r := range s {
		if !unicode.Is(unicode.Scripts["Han"], r) && "·" != string(r) {
			return false
		}
	}
	return true
}
