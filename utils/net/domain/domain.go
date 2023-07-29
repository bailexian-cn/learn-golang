package domain

import (
	"fmt"
	"regexp"
)

var domainRegexp = regexp.MustCompile("^[a-zA-Z0-9][-a-zA-Z0-9]{0,62}(\\.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+\\.?$")

func CheckDomainFormat(ip string) error {
	if !domainRegexp.MatchString(ip) {
		return fmt.Errorf("domain format error")
	}
	return nil
}
