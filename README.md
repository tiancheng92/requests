# requests
#### 发起http请求
##### 功能描述：
* 发起http请求。
##### 使用方法：
```go
response, err := requests.Request{
		Method: "", // 必填 请求方法 string 可选（"OPTIONS", "GET", "HEAD", "POST", "PUT", "DELETE", "TRACE", "CONNECT", "PATCH"）
		URL:    "", // 必填 请求地址 string
		Data:   "", // 非必填 请求包体(json) string
		Header: nil, // 非必填 请求头 map[string]string{}
	}.Run()
body := response.Body // 返回包体
status_code := response.StatusCode // 返回码
header := response.Header 
```

如有特殊需求请留言。