package __reflect

import (
	"fmt"
	"reflect"
	"testing"
)

type Interface interface {
	Name()
}

type CreateClusterCtl struct {

}

func (c CreateClusterCtl) Name() {

}

func TestGetObjTypeName(t *testing.T) {
	c := new(CreateClusterCtl)
	printName(c)

}

func printName(c Interface) {
	o := reflect.ValueOf(c)
	//obj := reflect.TypeOf(c.(struct{}))
	fmt.Println(o.Type())
	fmt.Println(o.Elem().Type().Name())
}
