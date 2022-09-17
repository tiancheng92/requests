package requests

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/tiancheng92/requests/json"
)

// request 发起请求的结构体
type request struct {
	Method     string
	URL        string
	Query      string
	Body       io.Reader
	Header     map[string]string
	Cookies    []*http.Cookie
	TimeOut    time.Duration
	TLS        *tls.Config
	ReadBody   bool
	FileFields map[string]string
	File       []*struct {
		FieldName string
		Filename  string
		FileData  io.Reader
	}
}

// New 新建一个Request对象
func New() *request {
	req := &request{ReadBody: true}
	return req
}

// GetRawResponseOnly 设置请求的url
func (r *request) GetRawResponseOnly() *request {
	r.ReadBody = false
	return r
}

// SetUrl 设置请求的url
func (r *request) SetUrl(url string) *request {
	r.URL = url
	return r
}

// SetRawBody 设置请求体（json或xml）
func (r *request) SetRawBody(data string) *request {
	r.Body = bytes.NewBuffer(strToBytes(data))
	return r
}

// SetJsonBody 设置Json请求体（结构体、Array、Dict...）
func (r *request) SetJsonBody(data any) *request {
	j, err := json.Marshal(data)
	if err != nil {
		panic("json marshal failed")
	}
	r.Body = bytes.NewBuffer(j)
	return r.AddJsonHeader()

}

// SetXMLBody 设置XML请求体（结构体、Array、Dict...）
func (r *request) SetXMLBody(data any) *request {
	j, err := xml.Marshal(data)
	if err != nil {
		panic("xml marshal failed")
	}
	r.Body = bytes.NewBuffer(j)
	return r.AddXMLHeader()
}

// SetRawQuery 设置Query（字符串 e.g:a=1&b=2）
func (r *request) SetRawQuery(data string) *request {
	r.Query = strings.TrimPrefix(data, "?")
	return r
}

// AddQuery 以k，v的形式逐一新增Query
func (r *request) AddQuery(key, value string) *request {
	ql := strings.Split(r.Query, "&")
	ql = append(ql, strings.Join([]string{key, value}, "="))
	if len(ql) > 0 {
		r.Query = strings.Join(ql, "&")
	}
	return r
}

// AddCookie 新增Cookie
func (r *request) AddCookie(cookie http.Cookie) *request {
	r.Cookies = append(r.Cookies, &cookie)
	return r
}

// AddHeader 新增Header头
func (r *request) AddHeader(key, value string) *request {
	if r.Header == nil {
		r.Header = make(map[string]string)
	}
	r.Header[key] = value
	return r
}

// AddFileField 新增上传文件时附带的key value对
func (r *request) AddFileField(key, value string) *request {
	if r.FileFields == nil {
		r.FileFields = make(map[string]string)
	}
	r.FileFields[key] = value
	return r
}

// AddFormHeader 添加表单请求头
func (r *request) AddFormHeader() *request {
	if r.Header == nil {
		r.Header = make(map[string]string)
	}
	r.Header["Content-Type"] = "application/x-www-form-urlencoded"
	return r
}

// AddJsonHeader 添加Json请求头
func (r *request) AddJsonHeader() *request {
	if r.Header == nil {
		r.Header = make(map[string]string)
	}
	r.Header["Content-Type"] = "application/json"
	return r
}

// AddXMLHeader 添加XML请求头
func (r *request) AddXMLHeader() *request {
	if r.Header == nil {
		r.Header = make(map[string]string)
	}
	r.Header["Content-Type"] = "text/xml"
	return r
}

// SetTimeOut 设置请求超时时间
func (r *request) SetTimeOut(duration time.Duration) *request {
	r.TimeOut = duration
	return r
}

// SetUploadFileByFilePath 设置上传的文件
func (r *request) SetUploadFileByFilePath(fieldName, filename string) *request {
	r.File = append(r.File, &struct {
		FieldName string
		Filename  string
		FileData  io.Reader
	}{fieldName, filename, nil})
	return r
}

func (r *request) SetUploadFile(fieldName, filename string, fileData io.Reader) *request {
	r.File = append(r.File, &struct {
		FieldName string
		Filename  string
		FileData  io.Reader
	}{fieldName, filename, fileData})
	return r
}

func (r *request) SetTLS(TlSConfig *tls.Config) *request {
	r.TLS = TlSConfig
	return r
}

func (r *request) SetBasicAuth(username, password string) *request {
	if r.Header == nil {
		r.Header = make(map[string]string)
	}
	r.Header["authorization"] = "Basic " + base64.StdEncoding.EncodeToString([]byte(username+":"+password))
	return r
}

// Get 发起Get请求
func (r *request) Get() (*Response, error) {
	r.Method = "GET"
	return r.run()
}

// Post 发起Post请求
func (r *request) Post() (*Response, error) {
	r.Method = "POST"
	return r.run()
}

// Put 发起Put请求
func (r *request) Put() (*Response, error) {
	r.Method = "PUT"
	return r.run()
}

// Patch 发起Patch请求
func (r *request) Patch() (*Response, error) {
	r.Method = "PATCH"
	return r.run()
}

// Delete 发起Delete请求
func (r *request) Delete() (*Response, error) {
	r.Method = "DELETE"
	return r.run()
}

// Head 发起Head请求
func (r *request) Head() (*Response, error) {
	r.Method = "HEAD"
	return r.run()
}

/************* 以下方法不对外暴露 **************/

// check 检测Request对象总的参数是否合法
func (r *request) check() error {
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

	// 判断文件上传所需的参数是否齐全
	for i := range r.File {
		if r.File[i] != nil {
			if r.File[i].Filename != "" && r.File[i].FieldName == "" {
				return errors.New("fieldname is empty")
			}
			if r.File[i].Filename == "" && r.File[i].FieldName != "" {
				return errors.New("filename is empty")
			}
		}
	}
	return nil
}

// getCompleteURL 获取完整的URL地址
func (r *request) getCompleteURL() string {
	if r.Query != "" {
		if strings.Contains(r.URL, "?") {
			return strings.Join([]string{r.URL, r.Query}, "&")
		}
		return strings.Join([]string{r.URL, r.Query}, "?")
	} else {
		return r.URL
	}
}

// setHeader 为http.Request对象设置请求头
func (r *request) setHeader(req *http.Request) *http.Request {
	if len(r.Header) > 0 {
		for k, v := range r.Header {
			req.Header.Set(k, v)
		}
	}
	return req
}

// setCookie 为http.Request对象设置Cookie
func (r *request) setCookie(req *http.Request) *http.Request {
	if len(r.Cookies) > 0 {
		for i := range r.Cookies {
			req.AddCookie(r.Cookies[i])
		}
	}
	return req
}

// setUploadBody 设置文件上传的请求体
func (r *request) setUploadBody() error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	for k, v := range r.FileFields {
		err := bodyWriter.WriteField(k, v)
		if err != nil {
			return err
		}
	}

	for i := range r.File {
		_, err := bodyWriter.CreateFormFile(r.File[i].FieldName, path.Base(r.File[i].Filename))
		if err != nil {
			return err
		}

		if r.File[i].FileData == nil {
			fb, err := os.ReadFile(path.Base(r.File[i].Filename))
			if err != nil {
				return err
			}
			bodyBuf.Write(fb)
		} else {
			_, err = io.Copy(bodyBuf, r.File[i].FileData)
			if err != nil {
				return err
			}
		}
	}
	contentType := bodyWriter.FormDataContentType()
	err := bodyWriter.Close()
	if err != nil {
		return err
	}
	r.Body = io.MultiReader(bodyBuf)
	r.AddHeader("Content-Type", contentType)
	return nil
}

// getBasisRequest 获取基础请求的http.Request对象
func (r *request) getBasisRequest() (*http.Request, error) {
	url := r.getCompleteURL()
	req, err := http.NewRequest(r.Method, url, r.Body)
	if err != nil {
		return nil, err
	}
	req = r.setHeader(req)
	req = r.setCookie(req)
	req.Close = true

	return req, nil
}

// getUploadRequest 获取支持文件上传的http.Request对象
func (r *request) getUploadRequest() (*http.Request, error) {
	url := r.getCompleteURL()
	err := r.setUploadBody()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(r.Method, url, r.Body)
	if err != nil {
		return nil, err
	}
	req = r.setHeader(req)
	req = r.setCookie(req)
	req.Close = true
	return req, nil
}

// run 执行request请求
func (r *request) run() (*Response, error) {
	var req *http.Request
	var err error

	// 检测数据
	if err = r.check(); err != nil {
		return nil, err
	}

	switch {
	case len(r.File) != 0:
		req, err = r.getUploadRequest()
	default:
		req, err = r.getBasisRequest()
	}
	if err != nil {
		return nil, err
	}

	// 开始请求
	client := &http.Client{Timeout: r.TimeOut, Transport: &http.Transport{TLSClientConfig: r.TLS}}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	// 读取内容
	var body []byte
	if r.ReadBody {
		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
	}

	return &Response{
		StatusCode:  resp.StatusCode,
		Body:        body,
		Header:      resp.Header,
		RawResponse: resp,
	}, nil
}
