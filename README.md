# requests
#### 发起http请求
##### 功能描述：
* 发起http请求。
##### 使用方法：
```go
package main

import (
	"fmt"
	"log"
	"github.com/tiancheng92/requests"
)

func ChainCall() {
	var r requests.Request
	res, err := r.SetUrl("").
		SetData("").
		AddHeader("", "").
		Post()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%#v\n", res)
}

func StructCall() {
	res, err := requests.Request{
		Method: "",
		URL:    "",
		Data:   "",
		Header: nil,
	}.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%#v", res)
}

func StructureCall() {
	var r requests.Request
	r.SetMethod("")
	r.SetUrl("")
	r.AddHeader("", "")
	// ...
	res, err := r.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%#v", res)
}
func main() {
	ChainCall()
	StructCall()
	StructureCall()
}

```
如有特殊需求请留言。