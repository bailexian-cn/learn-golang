defer语句是在一个函数在函数执行完毕之后，在更新返回值之后，`正常返回`、`出错返回`、`panic返回`时，在释放堆栈信息之前，执行的。

defer执行的时机

- 正常返回之前，可以修改返回值
- panic执行时，程序中断前，可用于打印错误堆栈信息

# 4.5.1 defer常见使用场景

## 成对操作

- 打开、关闭文件
- 连接、断开数据库
- 加锁、释放锁

示例：

```go
var mu sync.Mutex
var m = make(map[string]int)
func lookup(key string) int {
	mu.Lock()
	defer mu.Unlock()
	return m[key]
}
```

## 记录退出函数的时间

示例：

```go
func bigSlowOperation() {
	defer trace("bigSlowOperation")() // don't forget the extra parentheses
	// ...lots of work…
	time.Sleep(10 * time.Second) // simulate slow operation by sleeping
}
func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg,time.Since(start))
	}
}
```

defer语句中会先执行trace方法（记录开始时间）拿到函数值，等bigSlowOperation执行完成之后，再执行拿到的函数值（记录结束时间）。

## 修改返回值

因为defer语句中的函数会在更新返回值之后执行，所以defer语句中的函数可以用来更新返回值。

```go
func double(x int) (result int) {
	return x + x
}
func triple(x int) (result int) {
	defer func() { result += x }()
	return double(x)
}
fmt.Println(triple(4)) // "12"
```