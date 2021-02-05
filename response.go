package requests

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"net/http"
)

// Response request请求的返回结果
type Response struct {
	StatusCode int
	Body       ResponseBody
	Header     http.Header
}

type ResponseBody []byte

func (rb ResponseBody) JsonBind(obj interface{}) error {
	return json.NewDecoder(bytes.NewReader(rb)).Decode(obj)
}

func (rb ResponseBody) XMLBind(obj interface{}) error {
	return xml.NewDecoder(bytes.NewReader(rb)).Decode(obj)
}

func (rb ResponseBody) String() string {
	return string(rb)
}
