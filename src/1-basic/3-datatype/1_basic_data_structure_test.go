package __datatype

import (
	"fmt"
	"testing"
)

const (
	MaxInt32 = int32(-1) >> 1
	MinInt32 = int32(-1) << 31
)

func TestIntBit(t *testing.T) {
	a := 1<<62
	fmt.Println(a)

}
