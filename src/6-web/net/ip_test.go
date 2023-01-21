package net

import (
	"fmt"
	"testing"
)

func TestIpv4InCidr(t *testing.T) {
	cidr := "192.168.1.0/24"
	ip := "192.168.1.12"
	fmt.Println(Ipv4InCidr(ip, cidr))
}

func TestCidrConflict(t *testing.T) {
	cidr1 := "192.168.1.1/24"
	cidr2 := "192.168.0.0/16"
	fmt.Println(CidrConflict(cidr1, cidr2))
}
