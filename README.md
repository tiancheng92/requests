# requests

[![Build Status](https://github.com/tiancheng92/requests/workflows/Build/badge.svg)](https://github.com/tiancheng92/requests/actions)

### 功能描述：

* 发起http请求。

### 使用方法：

```go get github.com/tiancheng92/requests```

* example详解
  * 使用example前请优先启动./example/service服务
    * 该服务为一个mock的Http服务，用于相应request请求
  * ./example/basis_request 包含了一些基础的请求实例
  * ./example/upload_request 包含了文件上传的请求实例

#### 若要使用jsoniter进行json的序列化与放序列化请在构建参数中添加 `-tags=jsoniter`
#### 若要使用go_json进行json的序列化与放序列化请在构建参数中添加 `-tags=go_json`
