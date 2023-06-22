package commonpkg_fmt

import (
	"fmt"
	"testing"
)

func TestPrintf(t *testing.T) {
	fmt.Printf("%*s\n", 4, "a")
}
