package regexp

import (
	"fmt"
	"testing"
)

var validTestCases = []string{
	"192.168.122.1",
	"0.0.0.0",
}
var invalidTestCases = []string{
	"192.168.122",
	"192.168.122.x",
	"192.168.1?.2",
}

func TestIpv4Regex(t *testing.T) {
	fmt.Printf("start to test regexp %s\n", ipv4Regexp.String())
	for _, c := range validTestCases {
		if !ipv4Regexp.MatchString(c) {
			t.Errorf("case %s doesn't match doamin regexp\n", c)
		}
	}
	for _, c := range invalidTestCases {
		if ipv4Regexp.MatchString(c) {
			t.Errorf("case %s match doamin regexp\n", c)
		}
	}
}
