# 数组

**声明**

```go
var arr [3]int
arr1 := [3]int{1,2,3}
```

**访问**

```go
arr := [3]int{1,2,3}
a := a[2]
```

# 数组常用操作汇总

1. 获取长度
2. 获取容量
3. 从前查询
4. 从后查询
5. 扩容
6. 截取
7. 是否包含
8. 

# slice

切片是一个可变长的数组，其底层由数组实现包括长度、指针、容量三部分。

## **声明**

```go
var s []string
s1 := []string{"a"}
s2 := make([]string, 0, 5)  // 0表示容量，5表示长度
```

注意：slice的引用s1本身就是一个指针，数组也是一样。

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

## 从前查询



## **append**

```go
var s []string
s = append(s, "aa") // ["aa"]
s = append(s, s...) // ["aa","aa"]
s = append(s, s...) // ["aa","aa","aa","aa"]
```

append操作会新建一个对象将原有对象和需要扩展的对象放到新对象中并返回。

## **截取**

```go
s := []string{"a","b","c","d"}
s = s[:2] // ["a","b"]
```

截取返回的是一个**新的对象**

## **copy**

将slice前移一位，相当于将["d","e"]按序赋值给s[2:]

```go
s := []string{"a","b","c","d","e"}
copy(s[2:], s[3:]) // ["a", "b", "d", "e", "e"]
```

将slice后移一位，相当于将["c","d","e"]按序赋值给s[3:]，"e"超出容量就没赋值上去

```go
s := []string{"a","b","c","d","e"}
copy(s[3:], s[2:]) // ["a", "b", "c", "c", "d"]
```

