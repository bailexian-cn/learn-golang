# if

## 语法

```go
if condition {
	// statements
} else if condition {
    // statements
} else {
    // statements
}
```

# for

## 语法

```go
for initialization; condition; post {
	// statements
}
```

还有一种死循环的写法，可以通过break，return结束

```go
for {
	// statements
}
```

## for + judge

```golang
	arr := []int{1, 1, 1, 1, 0}
	i := 0
	for arr[i] > 0 {
		fmt.Printf("%d time\n", i)
		i++
	}
```

`arr[i] > 0`这个判断语句也可以是方法。

## 遍历slice

```go
arr := []string{"a", "b", "c"}
for i := range arr {
    fmt.Println(arr[i])
}
for index, value := range arr {
    fmt.Printf("index:%d; value:%s", index, value)
}
```

## 遍历map

```go
m := map[string]string{
    "a": "aa",
    "b": "bb",
}
for key := range m {
    fmt.Println(m[key])
}
```



# switch

简单示例

```go
var marks int = 90
switch marks {
    case 90: grade = "A"
    case 80: grade = "B"
    case 50,60,70 : grade = "C"
    default: grade = "D"  
}
```

注意：**Go语言并不需要显式地去在每一个case后写break**，语言默认执行完case后的逻辑语句会自动退出。如果你想要相邻的几个case都执行同一逻辑的话，需要自己显式地写上一个fallthrough语句来覆盖这种默认行为。

```go
t, c := 2, 2
var grade string
switch t {
    case 2: grade = "A";fallthrough
    case c: grade = "B"
    case 50,60,70 : grade = "C"
    default: grade = "D"
}
```

## type-switch

switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型。

示例：

```go
var x interface{}

switch i := x.(type) {
    case nil:  
    fmt.Printf(" x 的类型 :%T",i)                
    case int:  
    fmt.Printf("x 是 int 型")                      
    case float64:
    fmt.Printf("x 是 float64 型")          
    case func(int) int:
    fmt.Printf("x 是 func(int) int 型")                      
    case bool, string:
    fmt.Printf("x 是 abool 或 string 型" )      
    default:
    fmt.Printf("未知型")    
} 
```

# select

select 是 Go 中的一个控制结构，类似于用于通信的 switch 语句。每个 case 必须是一个通信操作，要么是发送要么是接收。

select 随机执行一个可运行的 case。如果没有 case 可运行，它将阻塞，直到有 case 可运行。一个默认的子句应该总是可运行的。

如果有 default 子句，则执行该语句。如果没有 default 子句，select 将阻塞，直到某个通信可以运行；Go 不会重新对 channel 或值进行求值。

## 语法

Go 编程语言中 select 语句的语法如下：

```golang
select {
  case communication clause  :
    statement(s);    
  case communication clause  :
    statement(s);
  /* 你可以定义任意数量的 case */
  default : /* 可选 */
    statement(s);
}
```

以下描述了 select 语句的语法：

- 每个 case 都必须是一个通信
- 所有 channel 表达式都会被求值
- 所有被发送的表达式都会被求值
- 如果任意某个通信可以进行，它就执行，其他被忽略。
- 如果有多个 case 都可以运行，Select 会随机公平地选出一个执行。其他不会执行。

## 实例

select 语句应用演示：

```golang
package main

import "fmt"

func main() {
  var c1, c2, c3 chan int
  var i1, i2 int
  select {
   case i1 = <-c1:
     fmt.Printf("received ", i1, " from c1\n")
   case c2 <- i2:
     fmt.Printf("sent ", i2, " to c2\n")
   case i3, ok := (<-c3): *// same as: i3, ok := <-c3*
     if ok {
      fmt.Printf("received ", i3, " from c3\n")
     } else {
      fmt.Printf("c3 is closed\n")
     }
   default:
     fmt.Printf("no communication\n")
  }   
}
```

以上代码执行结果为：

```
no communication
```