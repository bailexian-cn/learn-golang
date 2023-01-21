package __construction

import (
	"fmt"
	"testing"
)

type Cat struct {
	Name *string
}

// TestPointDeepCopy 指针的深拷贝示例
func TestPointDeepCopy(t *testing.T) {
	in := new(string)
	*in = "hello"
	out := new(string)
	*out = *in
	fmt.Printf("in: %p\n", in)
	fmt.Printf("out: %p\n", out)
}

// 指针属性结构体的深拷贝示例
func TestPointStructDeepCopy(t *testing.T) {
	b := "hello"
	in := &Cat{
		Name: &b,
	}
	out := new(Cat)
	*out = *in
	out.Name = new(string)
	*out.Name = *in.Name
	fmt.Printf("in: %+v\n", in)
	fmt.Printf("out: %+v\n", out)
	fmt.Printf("in.Name: %s\n", *in.Name)
	fmt.Printf("out.Name: %s\n", *out.Name)
}