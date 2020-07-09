package utils

import (
	"log"
	"net"
	"net/http"
	. "shadowDemo/zframework/logger"
	"shadowDemo/zframework/model"
	"strings"

	"github.com/mssola/user_agent"
	"github.com/sirupsen/logrus"
)

//获取真实ip
func GetRealIp(r *http.Request) string {
	value := r.Header.Get("X-Forwarded-For")
	if len(value) == 0 {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Printf("debug: Getting req.RemoteAddr %v", err)
			return ""
		}

		userIP := net.ParseIP(ip)
		if userIP == nil {
			log.Printf("debug: Parsing IP from Request.RemoteAddr got nothing.")
			return ""
		}
		return userIP.String()

	}
	Log.Debug("header X-Forwarded-For = %v", value)
	addresses := strings.Split(value, ",")
	address := strings.TrimSpace(addresses[0])
	return address
}

//获取硬件信息
func GetDeviceInfo(r *http.Request) model.Device {
	ua := user_agent.New(r.Header.Get("User-Agent"))
	browser, _ := ua.Browser()
	Log.WithFields(logrus.Fields{
		"ua":      FormatStruct(ua),
		"ua.os":   ua.OS(),
		"osinfo":  ua.OSInfo(),
		"browser": browser,
	}).Debug("getDeviceInfo")
	if ua.Mobile() {
		if strings.Contains(strings.ToLower(ua.OSInfo().Name), "iphone") {
			return model.IOS
		}
		if strings.Contains(strings.ToLower(ua.OSInfo().Name), "android") {
			return model.Android
		}
		return model.H5
	}
	return model.PC
}
