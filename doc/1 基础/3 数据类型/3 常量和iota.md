# 常量

## 声明

常量声明和变量声明一般都会出现在包级别，所以这些常量在整个包中都是可以共享的，或者你也可以把常量声明定义在函数体内部，那么这种常量就只能在函数体内用。目前常量声明的值必须是一个数字值、字符串或者一个固定的boolean值。

定义格式

```go
// 单个
const identifier [type] = value
// 多个
const c_name1, c_name2 = value1, value2

```

示例

```go
// 单个
const LENGTH int = 10
const WIDTH int = 5
// 多个，方式一
const a, b, c = 1, false, "str"
// 多个，方式二
const (
    Unknown = 1<<iota // 1
    Female =  // 2
    Male =   // 4
)
```

方式二中多个变量声明，除了第一个常量需要显示声明外，之后的声明都可以不显示写出来，继承之前常量的声明语句。

## iota

iota，特殊常量，可以认为是一个可以被编译器修改的常量。

iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。

示例1：

```go
package main

import "fmt"

func main() {
    const (
            a = iota   //0
            b          //1
            c          //2
            d = "ha"   //独立值，iota += 1
            e          //"ha"   iota += 1
            f = 100    //iota +=1
            g          //100  iota +=1
            h = iota   //7,恢复计数
            i          //8
    )
    fmt.Println(a,b,c,d,e,f,g,h,i)
}
/* 输出结果:
0 1 2 ha ha 100 100 7 8
*/
```

示例2：

```go
package main

import "fmt"
const (
    i=1<<iota
    j=3<<iota
    k
    l
)

func main() {
    fmt.Println("i=",i)
    fmt.Println("j=",j)
    fmt.Println("k=",k)
    fmt.Println("l=",l)
}
/* Output:
i= 1
j= 6
k= 12
l= 24
*/
```

iota 表示从 0 开始自动加 1，所以 **i=1<<0**, **j=3<<1**（**<<** 表示左移的意思），即：i=1, j=6，这没问题，关键在 k 和 l，从输出结果看 **k=3<<2**，**l=3<<3**。