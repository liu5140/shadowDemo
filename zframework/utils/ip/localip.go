package ip

import (
	"fmt"
	"log"
	"net"
)

// GET preferred outbound ip of this machine
func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println(localAddr.String())
	return localAddr.IP.String()
}

func GetLocalIP() (ip string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, addr := range addrs {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}
		return ipAddr.IP.String(), nil
	}
	return
}

// func GetLocalIp() (string, error) {
// 	addrs, err := net.InterfaceAddrs()
// 	var localIp string
// 	if err != nil {
// 		return localIp, err
// 	}
// 	for _, addr := range addrs {
// 		ip := addr.String()
// 		if (strings.Index(ip, "0.0.0.0") == -1) && (strings.Index(ip, "127.0.0.1") == -1) && (strings.Index(ip, ".") != -1) {
// 			localIp = strings.Split(ip, "/")[0] //fmt.Println("addr:", strings.Split(ip, "/")[0])
// 			break
// 		}
// 	}

// 	return localIp, nil
// }
