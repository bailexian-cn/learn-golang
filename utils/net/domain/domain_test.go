package domain

import "testing"

func TestCheckDomainFormat(t *testing.T) {
	validDomains := []string{
		"bailexian.cn",
		"minio.bailexian.cn",
	}
	invalidDomains := []string{
		"*33.baile.cnd.",
		"__fdkjflk",
	}
	for _, domain := range validDomains {
		e1 := CheckDomainFormat(domain)
		if e1 != nil {
			t.Errorf("except no err, but err %s for %s", e1, domain)
		}
	}
	for _, domain := range invalidDomains {
		e2 := CheckDomainFormat(domain)
		if e2 == nil {
			t.Errorf("except err, but no err for %s", domain)
		}
	}
}
