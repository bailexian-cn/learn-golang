# fmt.Sprintf

用于字符串拼接

```go
fmt.Sprintf("My name is %s", name)
```

## 字符串格式化

|      |                                                              |
| ---- | ------------------------------------------------------------ |
| %s   | string                                                       |
| %d   | 输出十进制                                                   |
| %b   | 输出二进制                                                   |
| %c   | 输出一个值的字符(char)                                       |
| %x   | 输出一个值的十六进制,每个字符串的字节用两个字符输出。打印的是一个值的16进制转型，例如int32(-1)输出的是-1，而不是0xffffffff |
| %f   | 输出浮点型数值                                               |
| %q   | 输出带双引号的字符串                                         |
| %p   | 输出一个指针的值                                             |
| %v   | 输出结构体的对象值                                           |
| %+v  | 输出结构体的成员名称和值                                     |
| %#v  | 输出一个值的Go语法表示方式                                   |
| %T   | 输出一个值的数据类型                                         |

## 对齐

|      |                                   |
| ---- | --------------------------------- |
| -5%s | 左对齐，长度小于5的字符串自动补齐 |
| 5%s  | 右对齐，长度小于5的字符串自动补齐 |
|      |                                   |

%*s

```golang
fmt.Printf("%*s\n", 4, "a") // 先打印4个空格，在打印a\n
```

# 打印结构体的内存地址

```golang
package main

import "fmt"

func main() {
	p1 := People{
		Name: "Lucy",
		Age: 32,
	}
	p1Copy := p1
	fmt.Println(fmt.Sprintf("p1 memory address is %p", &p1))
	fmt.Println(fmt.Sprintf("p1Copy memory address is %p", &p1Copy))
}

type People struct {
	Name string `json:"name,omitempty"`
	Age  uint8  `json:"age,omitempty"`
}
```



# tabwriter标准输出

func printTracks(tracks []*Track) {
const format = "%v\t%v\t%v\t%v\t%v\t\n"
tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
for _, t := range tracks {
fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
}
tw.Flush() // calculate column widths and print table
}