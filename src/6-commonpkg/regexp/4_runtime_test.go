package regexp

import (
	"fmt"
	"path"
	"runtime"
	"testing"
)

func TestGetCurrentDir(t *testing.T) {
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath := path.Dir(filename)
		fmt.Println(abPath)
	}
}
