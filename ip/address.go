package ip

import (
	"math"
	"strconv"
	"strings"
)

func sameIpAddress(ip1, ip2 string) bool {
	ip1_addr1, ip1_addr2 := ipAddress(ip1)
	ip2_addr1, ip2_addr2 := ipAddress(ip2)

	if ip1_addr1 == ip2_addr1 && math.Abs(float64(ip1_addr2-ip2_addr2)) <= 255 {
		return true
	}

	return false
}

func ipAddress(ip string) (addr1, addr2 int) {
	ips := strings.Split(ip, ".")
	if len(ips) != 4 {
		return
	}

	for k, v := range ips {
		i, err := strconv.Atoi(v)
		if err != nil || i > 255 {
			return
		}
		if k == 3 {
			addr2 = i
		} else {
			addr1 = addr1 | i<<uint(8*(3-k))
		}

	}
	return
}
