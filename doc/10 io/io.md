# IO

## os.Args

编译后执行文件及执行文件后参数组成的数组。

```go
// test.go
package main

import (
	"fmt"
	"os"
)
func main() {
	fmt.Println(os.Args)
}
```

执行

~> go run ./test.go aa bb cc

输出

[C:\Users\ADMINI~1\AppData\Local\Temp\go-build698832744\b001\exe\test.exe aa bb cc]

## os.Stdxxx

os里面一共有Stdin、Stdout、Stderr三个标准输入输出。

Stdin、Stdout和Stderr是打开的文件，指向标准输入、标准输出和标准错误文件描述符。

请注意，Go运行时写入用于panic和crashe的Stderr;关闭Stderr可能会导致这些消息转移到其他地方，可能是稍后打开的文件。

## os.File

文件描述符

### 判断文件是否存在

```go
func FileIsExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
```

### 打开文件

```
f, err := os.Open("/home/bailexian")
```

### 创建文件

```
f, err := os.Create(path)
```

### 删除文件

```
err := os.Remove(path)
```

### 读文件

```go
func ReadFile(absolutePath string) string {
	file, err := os.Open(absolutePath)

	if err != nil {
		fmt.Println("open file err: ", err)
		return err.Error()
	}
	// 读取文件缓冲区
	var buf [128]byte
	// 文件内容字节数组
	var content []byte
	for {
		n, err := file.Read(buf[:])
		if err == io.EOF {
			// 文件读取完毕
			break
		}

		if err != nil {
			fmt.Println("read file err: ", err)
			return err.Error()
		}

		content = append(content, buf[:n]...)
	}

	return string(content)
}
```

### 写文件

```
func WriteFile(absPath string, content string) {
	file, err := os.Create(absPath)
	if err != nil {
		fmt.Println("create file error: ", err)
		return
	}

	defer file.Close()

	file.WriteString(content)
}
```



## bufio.Scanner

bufio.Scanner是一个默认以行为分隔符的读取器，读取时每次读取输入的一行数据。

通过io.Reader创建一个Scanner

```
input := bufio.NewScanner(os.Stdin) // 标准输入，需要键入数据

file, err := os.Open(absolutePath)
fileInput := bufio.NewScanner(file) // 通过File创建Scanner，读取文件数据
```

Scanner对象通过Scan()方法可以读取一行数据，然后通过Text()方法将一行数据转为string，或者通过Bytes()方法将一行数据转为[]byte。

示例，将文件中的内容打印出来

```go
f, err := os.Open(absolutePath)
if err != nil {
    fmt.Errorf("\v\n", err)
    return
}
defer f.Close()
s := bufio.NewScanner(f)
for s.Scan() {
	fmt.Println(s.Text())
}
f.Close()
```

## ioutil.ReadFile(string)

将整个文件的内存读出来放到[]byte中

```go
ba, err := ioutil.ReadFile(path)
if err != nil {
fmt.Errorf("%v\n", err)
}
fmt.Println(string(ba))
```

## io.Copy(dst Writer, src Reader)

将输入src输出到dst中

返回值：written int64, err error

返回读取的数据大小

## ioutil.Discard

这是一个io.Writer

这是一个垃圾桶，可以向里面写一些不需要的数据