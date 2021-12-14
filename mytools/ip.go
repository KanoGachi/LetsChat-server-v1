package mytools

import (
	"net"
	"strings"
)

func GetOutBoundIP() (ip string) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return
	}
	ip = strings.Split(conn.LocalAddr().String(), ":")[0]
	return
}
