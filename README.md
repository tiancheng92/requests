[![Build Status](https://github.com/tiancheng92/requests/workflows/Build/badge.svg)](https://github.com/tiancheng92/requests/actions)

# 项目简介

* requests包是使用Go语言开发，模仿python的request包写的用于发起http请求，并返回响应结果的第三方包。

# 使用方法

```shell
go get -u github.com/tiancheng92/requests
```

# requests库方法详解

## requests（请求方法）

* 用户需要创建一个requests对象，然后调用相应的方法来设置请求体（链式调用），最后调用Get、Post、Put、Patch、Delete或Head方法发送请求。

| 函数签名                                                                                     | 描述                      |
|------------------------------------------------------------------------------------------|-------------------------|
| func New() *request                                                                      | 新建request对象             |
| func (r *request) SetUrl(url string) *request                                            | 设置请求的URL                |
| func (r *request) SetRawBody(data string) *request                                       | 设置请求体                   |
| func (r *request) SetJsonBody(data any) *request                                         | 设置JSON请求体并添加JSON请求头     |
| func (r *request) SetXMLBody(data any) *request                                          | 设置XML请求体并添加XML请求头       |
| func (r *request) SetRawQuery(data string) *request                                      | 设置Query（纯字符串，如：a=1&b=2） |
| func (r *request) AddQuery(key, value string) *request                                   | 以key/value对的形式逐一新增Query |
| func (r *request) AddCookie(cookie http.Cookie) *request                                 | 新增Cookie                |
| func (r *request) AddHeader(key, value string) *request                                  | 新增Header头               |
| func (r *request) AddFileField(key, value string) *request                               | 新增上传文件时附带的key/value对    |
| func (r *request) AddFormHeader() *request                                               | 添加表单请求头                 |
| func (r *request) AddJsonHeader() *request                                               | 添加JSON请求头               |
| func (r *request) AddXMLHeader() *request                                                | 添加XML请求头                |
| func (r *request) SetTimeOut(duration time.Duration) *request                            | 设置请求超时时间                |
| func (r *request) SetUploadFileByFilePath(fieldName, filename string) *request           | 设置要上传的文件路径              |
| func (r *request) SetUploadFile(fieldName, filename string, fileData io.Reader) *request | 设置要上传的文件                |
| func (r *request) SetTLS(TlSConfig *tls.Config) *request                                 | 设置TLS                   |
| func (r *request) SetBasicAuth(username, password string) *request                       | 设置Basic Auth            |
| func (r *request) Get() (*Response, error)                                               | 发起Get请求                 |
| func (r *request) Post() (*Response, error)                                              | 发起Post请求                |
| func (r *request) Put() (*Response, error)                                               | 发起Put请求                 |
| func (r *request) Patch() (*Response, error)                                             | 发起Patch请求               |
| func (r *request) Delete() (*Response, error)                                            | 发起Delete请求              |
| func (r *request) Head() (*Response, error)                                              | 发起Head请求                |

## response（返回内容）

* 成功请求远端数据后，会返回一个response对象，该对象包含了请求的状态码（Response.StatusCode）、请求头(StatusCode.Header)、响应体(Response.Body)
  、原生响应体（Response.RawResponse）。

### ResponseBody（响应体）

| 函数签名                                      | 描述                               |
|-------------------------------------------|----------------------------------|
| (rb ResponseBody) JsonBind(obj any) error | 把返回的JSON格式的Response.Body映射到指定对象中 |
| (rb ResponseBody) XMLBind(obj any) error  | 把返回的XML格式的Response.Body映射到指定对象中  |
| (rb ResponseBody) String() string         | 返回原生响应字符串                        |

# 构建相关

* requests库内建了自己的json包，支持使用[go_json](https://github.com/goccy/go-json)
  或[json-iterator](https://github.com/json-iterator/go)进行json解析。

| json解析方式                    | 构建约束                        |
|-----------------------------|-----------------------------|
| encoding/json               | go build ...                |
| github.com/goccy/go-json    | go build -tags=go_json ...  |
| github.com/json-iterator/go | go build -tags=jsoniter ... |