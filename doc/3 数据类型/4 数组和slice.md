数组是一个由固定长度的特定类型元素组成的序列，一个数组可以由零个或多个元素组成。

和数组对应的类型是Slice（切片），它是可以增长和收缩动态序列，slice功能也更灵活。

# 3.4.1 数组

```go
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
```

数组的大小是固定的，无法扩容。

# 3.4.2 slice

切片是一个可变长的数组，其底层由数组实现包括长度、指针、容量三部分。

## **声明**

```go
func TestDeclaration(t *testing.T) {
	var sli1 []int // 初始化值为nil
	sli2 := make([]int, 0, 0) // 第一个0表示长度，第二个0表示容量
	if sli1 == nil { // true
		fmt.Println("sli1's init val is nil")
	}
	if sli2 == nil { // false
		fmt.Println("sli2's init val is nil")
	}
    sli1 = append(sli1, 1) // nil值的slice也可扩容
	sli2 = append(sli2, 1) // 容量不够会自动扩容
	fmt.Printf("sli1's val is %v\n", sli1)
	fmt.Printf("sli2's val is %v\n", sli2)
}
```

**注意：slice的引用s1本身就是一个指针，数组也是一样。**

## **获取长度**

```go
var s []string
fmt.Println(len(s))
```

## **获取容量**

```go
var s []string
fmt.Println(cap(s))
```

## 扩容

```
func append(slice []Type, elems ...Type) []Type
```

```go
var s []string
s = append(s, "aa") // ["aa"]
s = append(s, s...) // ["aa","aa"]
s = append(s, s...) // ["aa","aa","aa","aa"]
```

append操作会**新建一个对象将原有对象和需要扩展的对象放到新对象中**并返回。

## 去重合并



## 缩容

```golang
func removeElement(slice []string, index int) []string {
    return append(slice[:index], slice[index+1:]...)
}
```

## **截取**

```go
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
/*
sli's memory address is        0xc0000122a0, val is [1 2 3 4]
cutOutSli's memory address is  0xc0000122a0, val is [1 2]
cutOutSli1's memory address is 0xc0000122b0, val is [3 4]
sli1's memory address is       0xc0000122a0, val is [1 2 3 4]
cutOutSli's val is [1 0]
*/
```

截取返回的是一个**新的对象**

## 拷贝

```
func copy(dst, src []Type) int
```

```go
func TestDeepCopy(t *testing.T) {
	sli := make([]int, 0)
	sli = append(sli, 1)
	//var copySli []int
	copySli := make([]int, len(sli))
	copy(copySli, sli)
	fmt.Printf("sli's memory address is         %p, val is %v\n", sli, sli)
	fmt.Printf("copySli's memory address is     %p, val is %v\n", copySli, copySli)
}
/*
sli's memory address is         0xc00000a348, val is [1]
copySli's memory address is     0xc00000a350, val is [1]
*/
```

拷贝实现前移、后移

```go
// 拷贝实现数组前移、后移
func TestDeepCopyMove(t *testing.T) {
   sli := []int{1, 2, 3, 4}
   fmt.Printf("sli's memory address is          %p, val is %v\n", sli, sli)
   copy(sli[1:], sli[2:]) // 最后两个元素向前移动1位，相当于将[3 4]按序赋值给sli[1:]，[1 3 4 4]
   fmt.Printf("forward sli's memory address is  %p, val is %v\n", sli, sli)
   copy(sli[1:3], sli[:2]) // 前面两个元素向后移动1位
   fmt.Printf("rearward sli's memory address is %p, val is %v\n", sli, sli)
}
```

## 查询

