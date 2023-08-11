# sdk-open-cloud-go

### 安装

```
go get github.com/w7corp/sdk-open-cloud-go
```

### 实例化sdk
```
client := w7.NewClient("appid", "apisecret")
```

### 更改sdk请求接口地址

````
client := w7.NewClient("appid", "apisecret", w7.Option{
    ApiUrl: "https://dev.w7.cc"
})
````
