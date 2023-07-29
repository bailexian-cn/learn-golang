package __datatype

import (
	"fmt"
	"strconv"
	"testing"
)

const (
	MaxInt32 = int32(-1) >> 1
	MinInt32 = int32(-1) << 31
)

func TestIntBit(t *testing.T) {
	a := 1 << 62
	fmt.Println(a)
}

// decimal convert to hexadecimal
func TestDec2Hex(t *testing.T) {
	i := 10
	hexStr := strconv.FormatInt(int64(i), 16)
	fmt.Println(hexStr)
}
