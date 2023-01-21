# 声明

```go
ages := make(map[string]int)
ages["alice"] = 31
ages["charlie"] = 34

ages1 := map[string]int{
    "alice": 31,
    "charlie": 34,
}
```

# 访问

```go
ages := map[string]int{
    "alice": 31,
    "charlie": 34,
}
if age, ok := ages["alice"]; ok {
    fmt.Printf("alice's age is %d", age)
}
```

如果访问的key在map中不存在，则会返回相应对象的0值。

禁止对map的元素作 取地址 操作，因为map可能随元素扩充重新分配内存，导致元素地址改变。

```go
_ = &ages["bob"] // compile error: cannot take address of map element
```



# 扩展

```go
ages := map[string]int{
    "alice": 31,
    "charlie": 34,
}
ages["darren"] = 26
```

# 删除

```go
ages := map[string]int{
    "alice": 31,
    "charlie": 34,
}
delete(ages, "alice")
```

# 遍历

```go
ages := map[string]int{
    "alice": 31,
    "charlie": 34,
}
for name,age := range ages {
	fmt.Printf("%s\t%d\n", name, age)
}
for ke := range ages {
    fmt.Printf("%s\t%d\n", ke, ages[ke])
}
```

对于map的遍历，取出元素的顺序每次都是随机的。每次都使用随机的遍历顺序可以强制要求程序不会依赖具体的哈希函数实现。

想要按顺序遍历，可以先出去keys，在对keys排序。

