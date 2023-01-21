package main

import "fmt"

func main() {
	forStmtSlice([]int{1, 2, 3})
}

func forStmtSlice(a []int) {
	for i := range a {
		fmt.Println(i)
	}
}