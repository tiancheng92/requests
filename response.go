package requests

import (
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
	return json.Unmarshal(rb, obj)
}

func (rb ResponseBody) XMLBind(obj interface{}) error {
	return xml.Unmarshal(rb, obj)
}

func (rb ResponseBody) String() string {
	return string(rb)
}
