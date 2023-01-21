# 正则匹配

匹配ipv4地址

```golang
ipv4Regexp = regexp.MustCompile("^((25[0-5])|(2[0-4]\\d)|(1\\d\\d)|([1-9]\\d)|\\d)(\\.((25[0-5])|(2[0-4]\\d)|(1\\d\\d)|([1-9]\\d)|\\d)){3}$")
ipv4Regexp.MatchString("192.168.122.1")
```

