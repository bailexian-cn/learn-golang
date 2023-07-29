package tcp

import (
	"fmt"
	"net"
	"regexp"
	"strconv"
)

// ipv4
var ipv4Regexp = regexp.MustCompile("^((25[0-5])|(2[0-4]\\d)|(1\\d\\d)|([1-9]\\d)|\\d)(\\.((25[0-5])|(2[0-4]\\d)|(1\\d\\d)|([1-9]\\d)|\\d)){3}$")

// cidr
var ipv4CidrRegexp = regexp.MustCompile("^(?:(?:[0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}(?:[0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\/([0-9]|[1-2]\\d|3[0-2])$")

func CheckIpv4Format(ip string) error {
	if !ipv4Regexp.MatchString(ip) {
		return fmt.Errorf("ipv4 format error")
	}
	return nil
}

func CheckIpv4CidrFormat(cidr string) error {
	if !ipv4CidrRegexp.MatchString(cidr) {
		return fmt.Errorf("ipv4 format error")
	}
	return nil
}

// ElectNoUsedIpFromSubnet 递增的方式从cidr中返回一个未使用的ipv4地址，函数会自动排除网络号、广播地址
// usedIpMap中的keys表示已经使用过的ipv4地址
func ElectNoUsedIpFromSubnet(cidr string, usedIpMap map[string]string) (string, error) {
	// elect vm cmd agent ip
	ip, cid, err := net.ParseCIDR(cidr)
	if err != nil {
		return "", err
	}
	netSize, _ := cid.Mask.Size()
	subSize := 32 - netSize
	max := uint(1<<subSize - 2) // 这里减2过滤掉广播地址
	var i uint
	ip = ip.To4()
	for i = 1; i <= max; i = i + 1 { // i从1开始过滤掉网络号
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
	return "", fmt.Errorf("all ipv4 addresses are used in %s", cidr)
}

// Ipv4InCidr ip是否在cidr内
// 不会校验输入，请使用者保证输入符合要求
func Ipv4InCidr(ip, cidr string) bool {
	_, cid, _ := net.ParseCIDR(cidr)
	mask, _ := cid.Mask.Size()
	ipn := net.ParseIP(ip)
	ipn = ipn.To4()
	ipint := int(ipn[0])<<24 + int(ipn[1])<<16 + int(ipn[2])<<8 + int(ipn[3])
	maskint := -1 >> (32 - mask) << (32 - mask)
	ipma := ipint & maskint
	ipman := net.IP{byte(ipma >> 24), byte(ipma >> 16), byte(ipma >> 8), byte(ipma)}
	return cid.IP.To4().Equal(ipman.To4())
}

// Ipv4CidrConflict 判断ipv4的cidr cidr1、cidr2是否冲突，如果冲突返回true，否则返回false
// 不会校验输入，请使用者保证输入符合要求，例如：cidr1，192.168.1.1/24；cidr2，192.168.0.0/16。
func Ipv4CidrConflict(cidr1, cidr2 string) bool {
	_, c1ipNet, _ := net.ParseCIDR(cidr1)
	_, c2ipNet, _ := net.ParseCIDR(cidr2)
	c1maskSize, _ := c1ipNet.Mask.Size()
	c2maskSize, _ := c2ipNet.Mask.Size()
	if c1maskSize < c2maskSize {
		cidr2 = c2ipNet.IP.To4().String() + "/" + strconv.Itoa(c1maskSize)
		_, c2ipNet, _ = net.ParseCIDR(cidr2)
	} else {
		cidr1 = c1ipNet.IP.To4().String() + "/" + strconv.Itoa(c2maskSize)
		_, c1ipNet, _ = net.ParseCIDR(cidr1)
	}

	return c1ipNet.IP.To4().Equal(c2ipNet.IP.To4())
}
