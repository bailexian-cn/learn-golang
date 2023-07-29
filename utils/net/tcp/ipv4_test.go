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
