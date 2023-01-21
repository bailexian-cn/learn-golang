package __exception

import (
	"fmt"
	"testing"
)

func TestRecover(t *testing.T) {

	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("c")
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容，55
		}
		fmt.Println("d")
	}()

	f()
}

func fp() {
	fmt.Println("a")
	panic(55)
	fmt.Println("b") // 这个语句无法执行
	fmt.Println("f")
}

func f() {
	fmt.Println("a")
	fmt.Println("b") // 这个语句无法执行
	fmt.Println("f")
}
