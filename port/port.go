package port

import (
	"net"
	"strconv"
	"time"
)

type ScanResult struct {
	Port  int
	State string
}

func ScanPort(protocol, hostname string, port int) ScanResult {
	result := ScanResult{Port: port}

	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)
	if err != nil {
		result.State = "Closed"
		return result
	}

	defer conn.Close()
	result.State = "Open"

	return result
}

func InitialScan(hostname string) []ScanResult {
	var results []ScanResult

	for i := 1; i <= 1024; i++ {
		result := ScanPort("tcp", hostname, i)
		results = append(results, result)
	}

	return results
}
