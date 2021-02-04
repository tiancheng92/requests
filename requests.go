package requests

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

// Request 发起请求的结构体
type Request struct {
	Method  string
	URL     string
	Body    io.Reader
	Query   string
	Header  map[string]string
	TimeOut time.Duration
	Cookies []*http.Cookie
	File    FileInfo
}

type FileInfo struct {
	Fieldname string
	Filename  string
}

func New() Request {
	return Request{}
}

func (r *Request) SetMethod(method string) *Request {
	r.Method = method
	return r
}

func (r *Request) SetUrl(url string) *Request {
	r.URL = url
	return r
}

func (r *Request) SetRawBody(data string) *Request {
	r.Body = bytes.NewBuffer([]byte(data))
	return r
}

func (r *Request) SetBody(data interface{}) *Request {
	j, err := json.Marshal(data)
	if err != nil {
		panic("json marshal failed")
	}
	r.Body = bytes.NewBuffer(j)
	return r
}

func (r *Request) SetRawQuery(data string) *Request {
	r.Query = strings.TrimPrefix(data, "?")
	return r
}

func (r *Request) AddQuery(key, value string) *Request {
	ql := strings.Split(r.Query, "&")
	ql = append(ql, fmt.Sprintf("%s=%s", key, value))
	if len(ql) > 0 {
		r.Query = strings.Join(ql, "&")
	}
	return r
}

func (r *Request) AddCookie(cookie http.Cookie) *Request {
	r.Cookies = append(r.Cookies, &cookie)
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

func (r *Request) UploadFile(fieldname, filename string) *Request {
	r.File = FileInfo{
		Fieldname: fieldname,
		Filename:  filename,
	}
	return r
}

func (r *Request) Get() (Response, error) {
	r.Method = "GET"
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

func (r *Request) Patch() (Response, error) {
	r.Method = "PATCH"
	return r.Run()
}

func (r *Request) check() error {
	// 判断方法是否合法
	var methodList = []string{"GET", "POST", "PUT", "DELETE", "PATCH"}
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
	if r.File.Filename != "" && r.File.Fieldname == "" {
		return errors.New("fieldname is empty")
	}
	if r.File.Filename == "" && r.File.Fieldname != "" {
		return errors.New("filename is empty")
	}
	return nil
}

func (r *Request) getHttpRequest() (*http.Request, error) {
	var url string

	if r.Query != "" {
		url = fmt.Sprintf("%s?%s", r.URL, r.Query)
	} else {
		url = r.URL
	}

	req, err := http.NewRequest(r.Method, url, r.Body)
	if err != nil {
		return nil, err
	}
	// 设置请求头,如无则添加默认请求头
	if len(r.Header) > 0 {
		for k, v := range r.Header {
			req.Header.Set(k, v)
		}
	} else {
		req.Header.Set("Content-Type", "application/json")
	}

	if len(r.Cookies) > 0 {
		for _, c := range r.Cookies {
			req.AddCookie(c)
		}
	}

	req.Close = true
	return req, nil
}

func (r *Request) getUploadRequest() (*http.Request, error) {
	body := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(body)

	fileWriter, err := bodyWriter.CreateFormFile(r.File.Fieldname, path.Base(r.File.Filename))
	if err != nil {
		return nil, err
	}

	f, err := os.Open(r.File.Filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	_, err = io.Copy(fileWriter, f)
	if err != nil {
		return nil, err
	}

	contentType := bodyWriter.FormDataContentType()
	err = bodyWriter.Close()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(r.Method, r.URL, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return req, nil
}

// Run 执行request请求
func (r Request) Run() (Response, error) {
	var req *http.Request
	var err error
	// 检测数据
	if err = r.check(); err != nil {
		return Response{}, err
	}
	switch {
	case r.File.Filename != "":
		req, err = r.getUploadRequest()
	default:
		req, err = r.getHttpRequest()
	}
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
		Body:       body,
		Header:     resp.Header,
	}, nil
}
