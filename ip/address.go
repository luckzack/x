package ip

import (
	"math"
	"strconv"
	"strings"
)

type IpAddress struct {
	Mask string
	Addr int
}

func SameIpAddressOfIP(ip1, ip2 string) bool {

	return SameIpAddress(GetIpAddress(ip1), GetIpAddress(ip2))
}

func SameIpAddress(addr1, addr2 *IpAddress) bool {
	if addr1.Mask == addr2.Mask && math.Abs(float64(addr1.Addr-addr2.Addr)) <= 255 {
		return true
	}
	return false
}

func GetIpAddress(ip string) *IpAddress {
	ips := strings.Split(ip, ".")
	if len(ips) != 4 {
		return nil
	}

	add := IpAddress{}

	add.Mask = strings.Join(ips[0:3], ".")
	add.Addr, _ = strconv.Atoi(ips[3])

	// for k, v := range ips {
	// 	i, err := strconv.Atoi(v)
	// 	if err != nil || i > 255 {
	// 		return nil
	// 	}
	// 	if k == 3 {
	// 		add.Addr = i
	// 	} else {
	// 		add.Mask = add.Mask | i<<uint(8*(3-k))
	// 	}

	// }
	return &add
}
