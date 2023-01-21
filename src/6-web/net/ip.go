package net

import (
	"fmt"
	"net"
	"strconv"
)

func ElectNoUsedIpFromSubnet(cidr string, usedIpMap map[string]string) (string, error) {
	// elect vm cmd agent ip
	ip, cid, err := net.ParseCIDR(cidr)
	if err != nil {
		return "", err
	}
	netSize, _ := cid.Mask.Size()
	subSize := 32 - netSize
	max := uint(1<<subSize - 2)
	var i uint
	ip = ip.To4()
	for i = 1; i <= max; i = i + 1 {
		tmpIp := net.IPv4(0, 0, 0, 0).To4()
		tmpIp[3] = ip[3] + (byte)(0xff&i)
		tmpIp[2] = ip[2] + (byte)(i>>8&0xff)
		tmpIp[1] = ip[1] + (byte)(i>>16&0xff)
		tmpIp[0] = ip[0] + (byte)(i>>24&0xff)
		tmpStr := tmpIp.String()
		if _, ok := usedIpMap[tmpStr]; !ok {
			return tmpStr, nil
		}
	}
	return "", fmt.Errorf("failed to elect vm cmd agent ip")
}

func Ipv4InCidr(ip, cidr string) bool {
	_, cid, _ := net.ParseCIDR(cidr)
	mask, _ := cid.Mask.Size()
	ipn := net.ParseIP(ip)
	ipn = ipn.To4()
	ipint := int(ipn[0]) << 24 + int(ipn[1]) << 16 + int(ipn[2]) << 8 + int(ipn[3])
	maskint := -1 >> (32-mask) << (32-mask)
	ipma := ipint & maskint
	ipman := net.IP{byte(ipma>>24), byte(ipma>>16), byte(ipma>>8), byte(ipma)}
	return cid.IP.To4().Equal(ipman.To4())
}

func CidrConflict(cidr1, cidr2 string) bool {
	_, c1ipNet, _ := net.ParseCIDR(cidr1)
	_, c2ipNet, _ := net.ParseCIDR(cidr2)
	c1maskSize,_  := c1ipNet.Mask.Size()
	c2maskSize,_  := c2ipNet.Mask.Size()
	if c1maskSize < c2maskSize {
		cidr2 = c2ipNet.IP.To4().String() + "/" + strconv.Itoa(c1maskSize)
		_, c2ipNet, _ = net.ParseCIDR(cidr2)
	} else {
		cidr1 = c1ipNet.IP.To4().String() + "/" + strconv.Itoa(c2maskSize)
		_, c1ipNet, _ = net.ParseCIDR(cidr1)
	}

	return c1ipNet.IP.To4().Equal(c2ipNet.IP.To4())
}
