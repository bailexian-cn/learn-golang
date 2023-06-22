package main

import (
	"fmt"
)

func main()  {
	fmt.Println("Hello")
}

func minInsertions(s string) int {
	arr := []rune(s)
	l := len(arr)
	var res, left int
	for i := 0; i < l; i++ {
		if arr[i] == '(' {
			left++
		} else {
			if left == 0 {
				res++
			} else {
				left--
			}
			if i == l -1 || arr[i+1] != ')' {
				res++
			} else {
				i++
			}
		}
	}
	return res + left * 2
}