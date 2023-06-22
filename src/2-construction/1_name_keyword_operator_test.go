package __construction

import (
	"fmt"
	"testing"
)

func TestBitOperate(t *testing.T) {
	i := -1
	fmt.Println(i << 1) // -2
	fmt.Println(i >> 1, uint(1) >> 1) // -1
	fmt.Println(i & 1)  // 1
	fmt.Println(i | 1)  // -1
	fmt.Println(^i)     // 0
}
