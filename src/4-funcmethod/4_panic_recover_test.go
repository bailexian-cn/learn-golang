package funcmethod

import (
	"fmt"
	"testing"
)

// 4.4.2 recover
func noReturn() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("internal error: %v\n", p)
		}
	}()
	panic("err")
}
func noReturn1() (res int) {
	defer func() {
		switch p := recover(); p {
		case nil:
			// no panic
		default:
			fmt.Println(p)
			res = 1
		}
	}()
	panic("err")
}
func soleTitle() (title string, err error) {
	type bailout struct{}
	defer func() {
		switch p := recover(); p {
		case nil:
			// no panic
		case bailout{}:
			// "expected" panic
			err = fmt.Errorf("multiple title elements")
		default:
			panic(p) // unexpected panic; carry on panicking
		}
	}()
	return title, nil
}
func TestRecover(t *testing.T) {
	noReturn()
	fmt.Println(noReturn1())
}
