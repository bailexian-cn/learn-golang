```golang
package main

import "fmt"

func main()  {
	err := Parse("HHH")
	fmt.Println(err)
}

func Parse(input string) (err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("internal error: %v", p)
		}
	}()
	panic(input)
}
```

可取的做法：

- 把panic的处理放在对应包下处理

web服务器将请求分发给处理函数，当遇到panic时会调用recover，输出堆栈信息，继续运行。可能引起资源泄露，或者导致其他问题。正确的做法是**有选择性的recover**。

有选择性的recover示例：

```

```

