package funcmethod

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

func paramsCheck(s string) error {
	if len(s) < 1 {
		return fmt.Errorf("param s can't be empty")
	}
	return nil
}

func Test1(t *testing.T) {
	if err := paramsCheck("a"); err != nil {
		fmt.Printf("err is %v", err)
	}
}

// 4.2.1 错误常见处理策略

// 4.2.2 文件结尾错误EOF
func Test2(t *testing.T) {
	in := strings.NewReader("haha")
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break // finished reading
		}
		fmt.Printf(string(r))
		if err != nil {
			fmt.Printf("read failed:%v\n", err)
		}
		// ...use r…
	}
	fmt.Println()
}
