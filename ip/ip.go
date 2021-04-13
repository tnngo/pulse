package ip

import (
	"errors"
	"net"
)

// GetLocalIP get local IP address.
func GetLocalIP() (string, error) {
	localAddrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, v := range localAddrs {
		if ipnet, ok := v.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", errors.New("Failed to obtain the local IP address")
}
