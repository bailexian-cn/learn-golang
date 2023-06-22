package funcmethod

import (
	"fmt"
	"testing"
)

// 4.3.1 函数值的类型
func TestFunValType(t *testing.T) {
	var fn func (i int) int
	fmt.Printf("fn's type is %T\n", fn) // func(int) int
	// fmt.Printf("%v", fn) // runtime err: fn is a func value, not called
	if fn == nil {
		fmt.Println("fn is nil now")
	}
	// fn(1) // panic
	//m := make(map[func (int) int]string) // compile err
}
// 4.3.2 函数值的局部变量
func TestFunValLocalVars(t *testing.T) {
	var prt func()
	fn1 := func() {
		var sum uint32
		sum++
		prt = func() {
			fmt.Println(sum)
		}
	}
	fn1()
	prt()
}
func TestFunValUsualErr(t *testing.T) {
	var prt func()
	var ar []func()
	fn1 := func() {
		var sum uint32
		for i := 0; i < 3; i++ {
			sum++
			a := sum
			prt = func() {
				fmt.Printf("%d ", a)
				//fmt.Printf("%d ", sum)
			}
			ar = append(ar, prt)
		}
	}
	fn1()
	for _, f := range ar {
		f()
	} // 3 3 3
}