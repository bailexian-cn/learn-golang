package __datatype

import (
	"fmt"
	"testing"
)

// 3.6.4 结构体的拷贝
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

func (in *Cat) DeepCopy1(out *Cat) {
	*out = *in
	if in.Name != nil {
		out.Name = new(string)
		*out.Name = *in.Name
	}
}

func (in *Cat) ErrDeepCopy(out *Cat) {
	*out = *in
	if in.Name != nil {
		in1, out1 := in.Name, out.Name
		out1 = new(string)
		*out1 = *in1
	}
}

func TestStructureCopy(t *testing.T) {
	catName := new(string)
	*catName = "Hellen"
	cat := &Cat{
		Name: catName,
		Age:  2,
	}
	copyCat := &Cat{}
	errCopyCat := &Cat{}
	cat.DeepCopy1(copyCat)
	cat.ErrDeepCopy(errCopyCat)
	fmt.Printf("cat's name memory address is                %p; name is %s\n", cat.Name, *cat.Name)
	fmt.Printf("copyCat's name memory address is            %p; name is %s\n", copyCat.Name, *copyCat.Name)
	fmt.Printf("errCopyCat's name memory address is         %p; name is %s\n", errCopyCat.Name, *errCopyCat.Name)
	fmt.Printf("cat's name pointer memory address is        %p; name is %s\n", &cat.Name, *cat.Name)
	fmt.Printf("copyCat's name pointer memory address is    %p; name is %s\n", &copyCat.Name, *copyCat.Name)
	fmt.Printf("errCopyCat's name pointer memory address is %p; name is %s\n", &errCopyCat.Name, *errCopyCat.Name)
}
