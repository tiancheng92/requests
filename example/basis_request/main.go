package main

import (
	"log"
	"requests"
	"time"
)

type mockData struct {
	Name string
	Date time.Time
}

// getRequests 发起get请求
func getRequests() {
	r := requests.New()
	res, err := r.SetUrl("http://127.0.0.1:8080/").
		AddQuery("key", "value").
		AddQuery("key", "value1").
		Get()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	log.Printf("%+v", res)      // {StatusCode:200 Body:method: GET, query: map[key:[value value1]] Header:map[Content-Length:[43] Content-Type:[text/plain; charset=utf-8] Date:[Fri, 05 Feb 2021 09:57:27 GMT]]}
	log.Printf("%+v", res.Body) // method: GET, query: map[key:[value value1]]
}

// postJsonRequests 发起post请求 包体为json字符串
func postJsonRequests() {
	r := requests.New()
	res, err := r.SetUrl("http://127.0.0.1:8080/").
		SetJsonBody(mockData{
			Name: "tiancheng92",
			Date: time.Now(),
		}).
		Post()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	log.Printf("%+v", res)      // {StatusCode:200 Body:{"Name":"tiancheng92","Date":"2021-02-05T16:51:26.615923+08:00"} Header:map[Content-Length:[64] Content-Type:[text/plain; charset=utf-8] Date:[Fri, 05 Feb 2021 08:47:56 GMT]]}
	log.Printf("%+v", res.Body) // {"Name":"tiancheng92","Date":"2021-02-05T16:51:26.615923+08:00"}
	var body mockData
	err = res.Body.JsonBind(&body)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	log.Printf("%s---%s", body.Name, body.Date.Format("2006-01-02 15:04:05")) // tiancheng92---2021-02-05 16:51:26
}

// postJsonRequests 发起post请求 包体为XML字符串
func postXMLRequests() {
	r := requests.New()
	res, err := r.SetUrl("http://127.0.0.1:8080/").
		SetXMLBody(mockData{
			Name: "tiancheng92",
			Date: time.Now(),
		}).
		Post()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	log.Printf("%+v", res)      // {StatusCode:200 Body:<mockData><Name>tiancheng92</Name><Date>2021-02-05T16:57:11.622405+08:00</Date></mockData> Header:map[Content-Length:[98] Content-Type:[text/plain; charset=utf-8] Date:[Fri, 05 Feb 2021 08:57:11 GMT]]}
	log.Printf("%+v", res.Body) // <mockData><Name>tiancheng92</Name><Date>2021-02-05T16:57:11.622405+08:00</Date></mockData>
	var body mockData
	err = res.Body.XMLBind(&body)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	log.Printf("%s---%s", body.Name, body.Date.Format("2006-01-02 15:04:05")) // tiancheng92---2021-02-05 16:57:11
}

// postRawRequests 发送原生的body
func postRawRequests() {
	r := requests.New()
	res, err := r.SetUrl("http://127.0.0.1:8080/").
		SetRawBody("<mockData><Name>tiancheng92</Name><Date>2021-02-05T16:57:11.622405+08:00</Date></mockData>").
		AddXMLHeader().
		Post()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	log.Printf("%+v", res)      // {StatusCode:200 Body:<mockData><Name>tiancheng92</Name><Date>2021-02-05T16:57:11.622405+08:00</Date></mockData> Header:map[Content-Length:[90] Content-Type:[text/plain; charset=utf-8] Date:[Fri, 05 Feb 2021 09:06:30 GMT]]}
	log.Printf("%+v", res.Body) // <mockData><Name>tiancheng92</Name><Date>2021-02-05T16:57:11.622405+08:00</Date></mockData>
	var body mockData
	err = res.Body.XMLBind(&body)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	log.Printf("%s---%s", body.Name, body.Date.Format("2006-01-02 15:04:05")) // tiancheng92---2021-02-05 16:57:11
}

// withTimeoutRequests 规定请求超时时间
func withTimeoutRequests() {
	r := requests.New()
	res, err := r.SetUrl("http://127.0.0.1:8080/with_time_out/").SetTimeOut(5 * time.Second).Get()
	if err != nil {
		log.Fatalf("%+v", err) // Get "http://127.0.0.1:8080/with_time_out/": context deadline exceeded (Client.Timeout exceeded while awaiting headers)
	}
	log.Printf("%+v", res)
	log.Printf("%+v", res.Body)
}

func main() {
	// 运行时请先开始example/service/mock_service服务
	getRequests()
	postJsonRequests()
	postXMLRequests()
	postRawRequests()
	withTimeoutRequests()
}
