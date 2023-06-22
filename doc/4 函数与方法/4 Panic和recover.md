# 4.4.1 panic介绍

有些错误只能在运行时检查，如数组访问越界、空指针引用等。这些运行时错误会引起painc异常。

一般而言，当panic异常发生时，程序会中断运行，所以panic一般用于严重的错误，可预见错误使用error机制。



# 4.4.2 recover

```golang
func noReturn() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("internal error: %v\n", p)
		}
	}()
	panic("err")
}
func TestRecover(t *testing.T) {
	noReturn()
}
```

