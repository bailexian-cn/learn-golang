package __datatype

import (
	"fmt"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	str := "hello word!"
	substr := "o"
	// 前序查询
	fmt.Println(strings.Index(str, substr)) // 4
	// 后序查询
	fmt.Println(strings.LastIndex(str, substr)) // 7
	// split
	fmt.Println(strings.Split(str, substr)) // [hell  w rd!]
	// contains
	fmt.Println(strings.Contains(str, substr)) // true
}

func TestTrimPrefix(t *testing.T) {
	str := "vip-lb-cke-xxxx"
	prefix := "vip-lb-"
	suffix := strings.TrimPrefix(str, prefix)
	if suffix != "cke-xxxx" {
		t.Errorf("except cke-xxxx")
	}
}
