JavaScript对象表示法（JSON）是一种用于发送和接收结构化信息的标准协议。

`示例一`：

```go
// 例一
package main

import (
	"encoding/json"
	"fmt"
)

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Circle struct {
	Point  `json:"point"`
	Radius int `json:"radius"`
}

func main() {
	c := Circle{
		Point: Point{
			X: 1,
			Y: 1,
		},
		Radius: 1,
	}
	s, _ := json.Marshal(c)    // 对象转json
	fmt.Println(string(s))     // {"point":{"x":1,"y":1},"radius":1}
	var c1 Circle
	_ = json.Unmarshal(s, &c1) // json转对象
	fmt.Println(c1.X)          // 1
    s1, _ := json.MarshalIndent(c, "", "  ")
}
```

# 对象转json

```go
s, err := json.Marshal(c)
```

对于对象c类型的定义，如果属性没有定义 `json:"x"`，那么输出json的属性名默认为成员名字，对于匿名成员如果没有写明json属性可就糟了，会直接把匿名成员的属性赋值给上级。例如`示例一`不加json属性的话输出的结果就为

```json
{"X":1,"Y":1,"Radius":1}
```

注意：**最好给所有结构体所有成员加上json属性**。

## **转json时，缩进**

```go
s1, err := json.MarshalIndent(c, "", "  ") // c是对象；第二个参数表示每一行输出的前缀，常用空字符串；第三个参数是每一个层级的缩进常用两个或者四个空格
```

上诉`示例一`的输出为

```json
{
  "point": {
    "x": 1,
    "y": 1
  },
  "radius": 1
}
```





```
1、只能包括的字符介绍
域名只能包括26个英文字母、“0-9”十个数字和“-.”英文中的连接号，但开头和结尾不能含有这个连接号。
英文域名：26个英文字母；“0”到“9”的数字；“-”英文中的连词不得用于开头及结尾处。
中文域名：两到十五个汉字之间的字词或词组；26个英文字母；“0”到“9”的数字。
```

