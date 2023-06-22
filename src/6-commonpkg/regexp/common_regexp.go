package regexp

import "regexp"

var domainRegexp, ipv4Regexp, ipv4CidrRegexp *regexp.Regexp

func init() {
	// 域名
	domainRegexp = regexp.MustCompile("^[a-zA-Z0-9][-a-zA-Z0-9]{0,62}(\\.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+\\.?$")
	// ipv4
	ipv4Regexp = regexp.MustCompile("^((25[0-5])|(2[0-4]\\d)|(1\\d\\d)|([1-9]\\d)|\\d)(\\.((25[0-5])|(2[0-4]\\d)|(1\\d\\d)|([1-9]\\d)|\\d)){3}$")
	// cidr
	ipv4CidrRegexp = regexp.MustCompile("^(?:(?:[0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}(?:[0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\/([0-9]|[1-2]\\d|3[0-2])$")
}
