# 获取当前位置

```golang
_, filename, _, ok := runtime.Caller(0)
if ok {
    abPath = path.Dir(filename)
}
```

