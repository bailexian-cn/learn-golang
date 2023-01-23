package __datatype

import (
	"fmt"
	"testing"
)

func TestStructureCopy(t *testing.T) {
	catName := new(string)
	*catName = "Hellen"
	cat := Cat{
		Name: catName,
		Age:  2,
	}
	catCopy := Cat{}
	cat.DeepCopy(&catCopy)
	fmt.Printf("cat's name memory address is %p\n", cat.Name)
	fmt.Printf("catCopy's name memory address is %p\n", catCopy.Name)
	fmt.Printf("cat's name pointer memory address is %p\n", &cat.Name)
	fmt.Printf("catCopy's name pointer memory address is %p\n", &catCopy.Name)
}

type Cat struct {
	Name *string `json:"name,omitempty"`
	Age  uint8   `json:"age,omitempty"`
}

func (in *Cat) DeepCopy(out *Cat) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}
