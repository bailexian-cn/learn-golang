go中可以通过`type`关键字来声明一个新的类型

# 声明规则

type 类型名字 底层类型

**示例**

```go
type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度
const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点温度
	BoilingC      Celsius = 100     // 沸水温度
)
```

# 新类型的比较

必须是两个相同类型的值才能进行比较。

底层类型支持的逻辑运算符可直接用于自定义类型。

```go
var c Celsius
var f Fahrenheit
fmt.Println(c == 0) // "true"
fmt.Println(f >= 0) // "true"
fmt.Println(c == f) // compile error: type mismatch
fmt.Println(c == Celsius(f)) // "true"!
```

# 用作枚举值

当我们需要创建一些枚举值，比如 虚拟机的状态时，可以使用type声明特定的类型。

例如

```golang
type EcsStatus string
const (
    Running=EcsStatus("running")
    /*...*/
)
```

