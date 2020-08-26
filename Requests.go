package requests

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// Request 发起请求的结构体
type Request struct {
	Method  string
	URL     string
	Data    string
	Header  map[string]string
	TimeOut time.Duration
}

// Response request请求的返回结果
type Response struct {
	StatusCode int
	Body       string
	Header     http.Header
}

func New() Request {
	return Request{}
}

func (res Response) JsonBind(obj interface{}) error {
	return json.NewDecoder(bytes.NewReader([]byte(res.Body))).Decode(obj)
}

func (r *Request) SetMethod(method string) *Request {
	r.Method = method
	return r
}

func (r *Request) SetUrl(url string) *Request {
	r.URL = url
	return r
}

func (r *Request) SetData(data string) *Request {
	r.Data = data
	return r
}

func (r *Request) SetStructData(data interface{}) *Request {
	j, err := json.Marshal(data)
	if err != nil {
		panic("json marshal failed")
	}
	r.Data = string(j)
	return r
}

func (r *Request) AddHeader(key, value string) *Request {
	if r.Header == nil {
		r.Header = make(map[string]string)
	}
	r.Header[key] = value
	return r
}

func (r *Request) SetTimeOut(duration time.Duration) *Request {
	r.TimeOut = duration
	return r
}

func (r *Request) Options() (Response, error) {
	r.Method = "OPTIONS"
	return r.Run()
}

func (r *Request) Get() (Response, error) {
	r.Method = "GET"
	return r.Run()
}

func (r *Request) Head() (Response, error) {
	r.Method = "HEAD"
	return r.Run()
}

func (r *Request) Post() (Response, error) {
	r.Method = "POST"
	return r.Run()
}

func (r *Request) Put() (Response, error) {
	r.Method = "PUT"
	return r.Run()
}

func (r *Request) Delete() (Response, error) {
	r.Method = "DELETE"
	return r.Run()
}

func (r *Request) Trace() (Response, error) {
	r.Method = "TRACE"
	return r.Run()
}

func (r *Request) Connect() (Response, error) {
	r.Method = "CONNECT"
	return r.Run()
}

func (r *Request) Patch() (Response, error) {
	r.Method = "PATCH"
	return r.Run()
}

func (r *Request) check() error {
	// 判断方法是否合法
	var methodList = []string{"OPTIONS", "GET", "HEAD", "POST", "PUT", "DELETE", "TRACE", "CONNECT", "PATCH"}
	var methodValid = false
	for i := 0; i < len(methodList); i++ {
		if methodList[i] == r.Method {
			methodValid = true
			break
		}
	}
	if !methodValid {
		return errors.New("invalid method")
	}
	// 判断URL是否为空
	if r.URL == "" {
		return errors.New("URL is empty")
	}
	// 判断有无请求包体，如有则判断其格式是否为json
	if r.Data != "" && !json.Valid([]byte(r.Data)) {
		return errors.New("invalid json")
	}
	return nil
}

func (r *Request) getHttpRequest() (*http.Request, error) {
	var data io.Reader

	if r.Data != "" {
		data = bytes.NewBuffer([]byte(r.Data))
	} else {
		data = nil
	}

	req, err := http.NewRequest(r.Method, r.URL, data)
	if err != nil {
		return nil, err
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
	return req, nil
}

// Run 执行request请求
func (r Request) Run() (Response, error) {
	// 检测数据
	if err := r.check(); err != nil {
		return Response{}, err
	}

	// 把数据绑定到http.request对象中
	req, err := r.getHttpRequest()
	if err != nil {
		return Response{}, err
	}

	// 开始请求
	client := &http.Client{Timeout: r.TimeOut}
	resp, err := client.Do(req)
	if err != nil {
		return Response{}, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	// 读取内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Response{}, err
	}

	return Response{
		StatusCode: resp.StatusCode,
		Body:       string(body),
		Header:     resp.Header,
	}, nil
}
