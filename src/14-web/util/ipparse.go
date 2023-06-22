package util

import (
	"fmt"
	"net"
)

func GetMaskStrFromCIDR(cidr string) (string,error) {
	if _, ipNet, err := net.ParseCIDR(cidr); err != nil {
		return "", err
	} else {
		mask := ipNet.Mask
		a, _ := mask.Size()
		fmt.Println(a)
		str := ""
		len := len(mask)
		for i, m := range mask{
			str += fmt.Sprintf("%d", m)
			if i < len - 1 {
				str += "."
			}
		}
		fmt.Println(mask)
		return string(ipNet.Mask), nil
	}
}
