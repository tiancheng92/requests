package requests

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/wxnacy/wgo/arrays"
)

// Request 发起请求的结构体
type Request struct {
	Method string
	URL    string
	Data   string
	Header map[string]string
}

// Response request请求的返回结果
type Response struct {
	StatusCode int
	Body       string
	Header     http.Header
}

func (r *Request) SetMethod(method string) {
	r.Method = method
}

func (r *Request) SetUrl(url string) {
	r.URL = url
}

func (r *Request) SetData(data string) {
	r.Data = data
}

func (r *Request) AddHeader(key, value string) {
	r.Header[key] = value
}

func (r *Request) Options() (res Response, err error) {
	r.Method = "OPTIONS"
	return r.Run()
}

func (r *Request) Get() (res Response, err error) {
	r.Method = "GET"
	return r.Run()
}

func (r *Request) Head() (res Response, err error) {
	r.Method = "HEAD"
	return r.Run()
}

func (r *Request) Post() (res Response, err error) {
	r.Method = "POST"
	return r.Run()
}

func (r *Request) Put() (res Response, err error) {
	r.Method = "PUT"
	return r.Run()
}

func (r *Request) Delete() (res Response, err error) {
	r.Method = "DELETE"
	return r.Run()
}

func (r *Request) Trace() (res Response, err error) {
	r.Method = "TRACE"
	return r.Run()
}

func (r *Request) Connect() (res Response, err error) {
	r.Method = "CONNECT"
	return r.Run()
}

func (r *Request) Patch() (res Response, err error) {
	r.Method = "PATCH"
	return r.Run()
}

// Run 执行request请求
func (r *Request) Run() (res Response, err error) {
	var req *http.Request
	// 判断方法是否合法
	if arrays.Contains([]string{"OPTIONS", "GET", "HEAD", "POST", "PUT", "DELETE", "TRACE", "CONNECT", "PATCH"}, r.Method) < 0 {
		return Response{}, errors.New("invalid method")
	}
	// 判断URL是否为空
	if r.URL == "" {
		return Response{}, errors.New("URL is empty")
	}
	// 判断有无请求包体，如有则判断其格式是否为json
	if r.Data != "" {
		if !json.Valid([]byte(r.Data)) {
			return Response{}, errors.New("invalid json")
		}
		req, err = http.NewRequest(r.Method, r.URL, bytes.NewBuffer([]byte(r.Data)))
	} else {
		req, err = http.NewRequest(r.Method, r.URL, nil)
	}
	if err != nil {
		return
	}
	// 设置请求头,如无则添加默认请求头
	if len(r.Header) > 0 {
		for k, v := range r.Header {
			req.Header.Set(k, v)
		}
	} else {
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	}
	req.Close = true
	// 开始请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	// 读取内容
	body, _ := ioutil.ReadAll(resp.Body)
	res.StatusCode = resp.StatusCode
	res.Body = string(body)
	res.Header = resp.Header
	return
}
