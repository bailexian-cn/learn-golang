package __commonpkg

import (
	"net"
	"testing"
)

func TestParseIPv4(t *testing.T) {
	//a := net.IP("192.168.0.1")
	//if a.String() != "192.168.0.1" {
	//	t.Errorf("ip parse error")
	//}
	b := net.ParseIP("192.168.0.1")
	if b.String() != "192.168.0.1" {
		t.Errorf("ip parse error")
	}
}
