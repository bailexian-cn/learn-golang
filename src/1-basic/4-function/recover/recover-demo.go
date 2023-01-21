package main

import "fmt"

func main()  {
	err := Parse("HHH")
	fmt.Println(err)
}

func Parse(input string) (err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("internal error: %v", p)
		}
	}()
	panic(input)
}
