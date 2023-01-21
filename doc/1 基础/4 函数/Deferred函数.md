# 介绍

defer语句是在一个函数在函数执行完毕之后，在更新返回值之后，`正常返回`、`出错返回`、`panic返回`时，在释放堆栈信息之前，执行的。

注意：defer后接的语句会按顺序执行，defer后接的函数会在函数执行完毕后执行。可参考`记录退出函数的时间`中的示例`defer trace("bigSlowOperation")()`。

# 使用场景

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

bigSlowOperation函数中会先执行trace方法，当bigSlowOperation函数执行完成后执行defer指定的方法，此方法是trace方法的返回值。

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

# 常见问题

## 局部变量

下面的代码会导致系统的文件描述符耗尽，因为在所有文件都被处理之前，没有文件会被关闭。

```go
for _, filename := range filenames {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close() // NOTE: risky; could run out of file descriptors
	// ...process f…
}
```

