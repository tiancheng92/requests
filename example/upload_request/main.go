package main

import (
	"log"
	"requests"
)

func uploadRequest() {
	r := requests.New()
	res, err := r.SetUrl("http://127.0.0.1:8080/upload/").SetUploadFile("uploadfile", "./upload_file.txt").Post()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	log.Printf("%+v", res)      // {StatusCode:200 Body:map[Content-Disposition:[form-data; name="uploadfile"; filename="upload_file.txt"] Content-Type:[application/octet-stream]]success Header:map[Content-Length:[130] Content-Type:[text/plain; charset=utf-8] Date:[Fri, 05 Feb 2021 09:47:25 GMT]]}
	log.Printf("%+v", res.Body) // map[Content-Disposition:[form-data; name="uploadfile"; filename="upload_file.txt"] Content-Type:[application/octet-stream]]success
}

func main() {
	// 运行时请先开始example/service/mock_service服务
	uploadRequest()
}
