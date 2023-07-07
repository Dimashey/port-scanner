package port

import (
	"net"
	"strconv"
	"time"
)

func ScanPort(protocol, hostname string, port int) bool {
	address := hostname + ":" + strconv.Itoa(port)

	connection, err := net.DialTimeout(protocol, address, 60*time.Second)
	if err != nil {
		return false
	}

	defer connection.Close()

	return true
}
