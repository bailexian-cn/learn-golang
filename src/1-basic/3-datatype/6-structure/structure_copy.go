package main

import "fmt"

func main() {
	catName := new(string)
	*catName = "Hellen"
	cat := Cat {
		Name: catName,
		Age: 2,
	}
	catCopy := Cat{}
	cat.DeepCopy(&catCopy)
	fmt.Println(fmt.Sprintf("cat's name memory address is %p", cat.Name))
	fmt.Println(fmt.Sprintf("catCopy's name memory address is %p", catCopy.Name))
	fmt.Println(fmt.Sprintf("cat's name pointer memory address is %p", &cat.Name))
	fmt.Println(fmt.Sprintf("catCopy's name pointer memory address is %p", &catCopy.Name))
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

