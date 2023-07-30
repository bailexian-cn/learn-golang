package tcp

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
	fmt.Println(Ipv4CidrConflict(cidr1, cidr2))
}

func TestElectNoUsedIpFromSubnet(t *testing.T) {
	cidr := "192.168.0.0/16"
	usedIpMap := map[string]string{
		"192.168.0.3": "",
		"192.168.1.3": "",
	}
	noUsedIp, err := ElectNoUsedIpFromSubnet(cidr, usedIpMap)
	if err != nil {
		t.Error(err)
	}
	if noUsedIp != "192.168.0.1" {
		t.Errorf("except 192.168.0.0, but %s", noUsedIp)
	}
}

func TestCheckIpv4Format(t *testing.T) {
	validIps := []string{
		"192.168.0.1",
		"1.1.1.1",
	}
	invalidIps := []string{
		"1.1.1.1.",
		"a.1.3",
	}
	for _, ip := range validIps {
		e1 := CheckIpv4Format(ip)
		if e1 != nil {
			t.Error(e1)
		}
	}
	for _, ip := range invalidIps {
		e2 := CheckIpv4Format(ip)
		if e2 == nil {
			t.Errorf("except err, but no err for %s", ip)
		}
	}
}

func TestCheckIpv4CidrFormat(t *testing.T) {
	validCidrs := []string{
		"192.168.0.0/16",
		"1.1.1.1/24",
	}
	invalidCidrs := []string{
		"1.1.1.3",
		"3.3.3.3/44",
	}
	for _, cidr := range validCidrs {
		e1 := CheckIpv4CidrFormat(cidr)
		if e1 != nil {
			t.Errorf("execpt no err, but err %s for %s", e1, cidr)
		}
	}
	for _, cidr := range invalidCidrs {
		e2 := CheckIpv4CidrFormat(cidr)
		if e2 == nil {
			t.Errorf("except err, but no err for %s", cidr)
		}
	}
}
