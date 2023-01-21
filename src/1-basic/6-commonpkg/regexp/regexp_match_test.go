package regexp

import (
	"regexp"
	"testing"
)

func TestRegexpMatch(t *testing.T) {
	reg := regexp.MustCompile("[a-z0-9]([-a-z0-9]*[a-z0-9])?(.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*")
	res := reg.Match([]byte("ceake-demo"))
	if !res {
		t.Errorf("can't match regexp")
	}
}

func TestDomainRegexp(t *testing.T) {
	validTestCases := [][]byte{
		[]byte("ceake.com"),
	}
	invalidTestCases := [][]byte{
		[]byte("ceake.com&."),
	}
	for _, c := range validTestCases {
		if !domainRegexp.Match(c) {
			t.Errorf("case %s doesn't match doamin regexp\n", c)
		}
	}
	for _, c := range invalidTestCases {
		if domainRegexp.Match(c) {
			t.Errorf("case %s match doamin regexp\n", c)
		}
	}
}
