# 示例：将<库包>import到vendor中

将库包文件code.cestc.cn/ccos/envoy-openapi-go-sdk拷贝到项目vendor目录下

在go.mod文件中加入下列语句

```
require code.cestc.cn/ccos/envoy-openapi-go-sdk v0.0.0-00010101000000-000000000000
replace code.cestc.cn/ccos/envoy-openapi-go-sdk => /root/code/envoy-openapi-go-sdk
```

