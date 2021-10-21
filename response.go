package requests

import (
	"encoding/xml"
	"net/http"

	"github.com/tiancheng92/requests/json"
)

// Response request请求的返回结果
type Response struct {
	StatusCode  int
	Body        ResponseBody
	Header      http.Header
	RawResponse *http.Response
}

type ResponseBody []byte

func (rb ResponseBody) JsonBind(obj interface{}) error {
	return json.Unmarshal(rb, obj)
}

func (rb ResponseBody) XMLBind(obj interface{}) error {
	return xml.Unmarshal(rb, obj)
}

func (rb ResponseBody) String() string {
	return bytesToStr(rb)
}
