package __datatype

import (
	"fmt"
	"testing"
)

type Node struct {
	Name string
	Ip   string
}

// 3.4.1 使用数组
func Test1(t *testing.T) {
	var arr [3]int
	arr1 := [3]int{1,2,3}
	// arr2 := append(arr1, 4) // compile err
	fmt.Println(arr)
	fmt.Printf("%T\n", arr1) // [3]int
	sli := make([]int, 0)
	// sli = arr1 // compile err
	sli = append(sli, 1)
	fmt.Printf("%T\n", sli) // []int
}

// 3.4.2 slice
// 声明
func TestDeclaration(t *testing.T) {
	var sli1 []int
	sli2 := make([]int, 0, 0)
	if sli1 == nil { // true
		fmt.Println("sli1's init val is nil")
	}
	if sli2 == nil { // false
		fmt.Println("sli2's init val is nil")
	}
	sli1 = append(sli1, 1) // nil值的slice也可扩容
	sli2 = append(sli2, 1)
	fmt.Printf("sli1's val is %v\n", sli1)
	fmt.Printf("sli2's val is %v\n", sli2)
}
// 截取
func TestCutOut(t *testing.T) {
	sli := []int{1, 2, 3, 4}
	cutOutSli := sli[:2] // 没有深拷贝,不包含sli[2] [1 2]
	cutOutSli1 := sli[2:]// 包含sli[2] [3 4]
	sli1 := sli[:]       // 没有深拷贝 [1 2 3 4]
	fmt.Printf("sli's memory address is        %p, val is %v\n", sli, sli)
	fmt.Printf("cutOutSli's memory address is  %p, val is %v\n", cutOutSli, cutOutSli)
	fmt.Printf("cutOutSli1's memory address is %p, val is %v\n", cutOutSli1, cutOutSli1)
	fmt.Printf("sli1's memory address is       %p, val is %v\n", sli1, sli1)
	// change sli[1]
	sli[1] = 0
	fmt.Printf("cutOutSli's val is %v\n", cutOutSli)
}
// 拷贝
func TestDeepCopy(t *testing.T) {
	sli := make([]int, 0)
	sli = append(sli, 1)
	//var copySli []int
	copySli := make([]int, len(sli))
	copy(copySli, sli)
	fmt.Printf("sli's memory address is         %p, val is %v\n", sli, sli)
	fmt.Printf("copySli's memory address is     %p, val is %v\n", copySli, copySli)
}
// 拷贝实现数组前移、后移
func TestDeepCopyMove(t *testing.T) {
	sli := []int{1, 2, 3, 4}
	fmt.Printf("sli's memory address is          %p, val is %v\n", sli, sli)
	copy(sli[1:], sli[2:]) // 最后两个元素向前移动1位
	fmt.Printf("forward sli's memory address is  %p, val is %v\n", sli, sli)
	copy(sli[1:3], sli[:2]) // 前面两个元素向后移动1位
	fmt.Printf("rearward sli's memory address is %p, val is %v\n", sli, sli)
}

func TestChangePointArr(t *testing.T) {
	var pointerArr = []*Node{
		{
			Name: "node1",
			Ip:   "192.168.122.1",
		},
		{
			Name: "node2",
			Ip:   "192.168.122.2",
		},
	}
	newIp := "xx.xx.xx.xx"
	name := "node1"
	for _, n := range pointerArr {
		if n.Name == name {
			n.Ip = newIp
		}
	}
	fmt.Printf("pointerArr after changed is:\n")
	for _, n := range pointerArr {
		if n.Name == name && n.Ip != newIp {
			t.Fatalf("except: { \"name\": \"node1\", \"ip\": \"xx.xx.xx.xx\"}")
		}
	}
}

func TestChangeObjectArr(t *testing.T) {
	var objectArr = []Node{
		{
			Name: "node1",
			Ip:   "192.168.122.1",
		},
		{
			Name: "node2",
			Ip:   "192.168.122.2",
		},
	}
	newIp := "xx.xx.xx.xx"
	name := "node1"
	for _, n := range objectArr {
		if n.Name == name {
			n.Ip = newIp
		}
	}
	fmt.Printf("pointerArr after changed is:\n")
	for _, n := range objectArr {
		if n.Name == name && n.Ip == newIp {
			t.Fatalf("except: { \"name\": \"node1\", \"ip\": \"192.168.122.1\"}")
		}
	}
}

func TestArrIndex(t *testing.T) {
	arr := []string{"aa", "bb", "cc", "aa"}
	fmt.Printf("arr: %v\n", arr)
}
