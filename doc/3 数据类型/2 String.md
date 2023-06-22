字符串就是一串固定长度的字符连接起来的字符序列。Go 的字符串是由单个字节连接起来的。Go 语言的字符串的字节使用 UTF-8 编码标识 Unicode 文本。

# 声明

```go
a := "愿君多采撷，\n此物最相思。"
b := 'a'
```

单引号只能赋值**最多一个字符**

**不转意声明**

```go
	a := `愿君多采撷，
此物最相思。`
```

`，`后面的回车符会一并赋值给变量

# 转型

## To string

### 整型 To string

```go
i := 25105
s := strconv.Itoa(i) // 25105
s1 := strconv.FormatInt(int64(i), 10) // 25105，转为10进制字符串
```

### 浮点型 To string

```go
i := 25105.4342
s := strconv.FormatFloat(i, 'f', 10, 64) // 25105.4342000000，64位机器将i精确到小数点10位，f表示不做科学计数
s1 := strconv.FormatFloat(i, 'e', 10, 64) // 2.5105434200e+04，64位机器将i精确到小数点10位，e表示以10为幂做科学计数
```

### slice To string

```go
strings.Join([]string,string)
```

示例：

```golang
// 参数用空格隔开
strings.Join(os.Args[1:], " ")
```

### []byte To string

默认使用UTF-8编码

```go
arr := []byte{72,101,108,108,111,32,87,111,114,108,100}
str := string(arr) // Hello World
```



## string To

### string To 整型

```go
s := "fff"
i, _ := strconv.ParseInt(s, 16, 64) // 4095
```

### string To 浮点型

```go
s := "2.5553e+04"
i, _ := strconv.ParseFloat(s, 64) // 25553
```

### string To []byte

```go
str := "我你"
arr := []byte(str)
```

### string To slice

```golang
s := "aaa\nbbb\nccc"
arr := strings.Split(s,"\n")
```

# 字符串常用操作汇总

1. 获取长度
2. 从前查询，index()
3. 从后查询，lastIndex()
4. 分割，split()
5. 是否包含，contains()
6. 拼接
7. 截取
8. 以xxx开始

# 字符串操作

## 获取长度

```golang
str := "我good"
fmt.Println(len([]rune(str))) // 字符长度：5
fmt.Println(len([]byte(str))) // 字节长度：8
```

## 查询

```
str := "hello word!"
substr := "o"
fmt.Println(strings.Index(str, substr)) // 4
```

## 从后查询

```golang
str := "hello word!"
substr := "o"
fmt.Println(strings.LastIndex(str, substr)) // 4
```

## 分割

```
str := "hello word!"
substr := "o"
fmt.Println(strings.Split(str, substr)) // [hell  w rd!]
```

## 是否包含

```go
strings.Contains("沃尼塔", "我") // false
strings.Contains("沃尼塔", "沃") // true
```

## 拼接

```golang
fmt.Sprintf("imageExporterNodeSelector.%s", "role/minio")
```

## 截取

```go
s := "abcdef"
s1 := s[:3] // abc，截取前3个字节，注意是字节
s2 := "我你他" // [230 136 145 228 189 160 228 187 150]
s3 := s2[:3] // 我，UTF-8每个中文字符占三个字节
```

截取字符

```go
s2 := "我fdd你d他" 
s3 := string([]rune(s2)[:2]) // 我f
```

## 以xxx开始

```golang
strings.HasPrefix(url, startUrl)
```

