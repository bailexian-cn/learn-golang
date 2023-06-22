package funcmethod

import (
	"fmt"
	"testing"
)

// 4.1.1 声明，示例
// count compute char count of a and b
func StringCharLenSum(a, b string) int {
	return len([]rune(a)) + len([]rune(b))
}

func StringByteLenSum(a, b string) int {
	return len([]byte(a)) + len([]byte(b))
}

func TestStringCharLenSum(t *testing.T) {
	if StringCharLenSum("我hfd发", "哈") != 6 {
		t.Fatal("except char count sum is 6")
	}
}

func TestStringByteLenSum(t *testing.T) {
	if StringByteLenSum("我hfd发", "哈") != 12 {
		t.Fatal("except byte count sum is 12")
	}
}

// 4.1.2 可变入参，示例
func printArr(arr ...int) {
	for _, a := range arr {
		fmt.Printf("%d ", a)
	}
	fmt.Println()
}

func TestVariableParameter(t *testing.T) {
	printArr(1, 2, 3)
	printArr([]int{1, 2, 3}...)
}

// 4.1.3 多返回值，示例
func parse(s string) (string, error) {
	return s, nil
}

func parse1() (string, error) {
	return parse("hello world!")
}

func printParseRes(s string, err error) {
	fmt.Printf("parse(%s, %v) res is: ...\n", s, err)
}

func TestSeveralReturnVals(t *testing.T) {
	printParseRes(parse1())
}