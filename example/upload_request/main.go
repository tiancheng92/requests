package main

import (
	"log"
	"os"

	"github.com/tiancheng92/requests"
)

func uploadRequest() {
	r := requests.New()
	f, err := os.Open("./upload_file.txt")
	if err != nil {
		log.Fatalf("%+v", err)
		return
	}
	defer f.Close()
	res, err := r.SetUrl("http://127.0.0.1:8080/upload/").
		SetUploadFileByFilePath("uploadfile", "./upload_file.txt").
		Post()
	if err != nil {
		log.Fatalf("%+v", err)
		return
	}

	log.Printf("%+v", res)      // {StatusCode:200 Body:map[Content-Disposition:[form-data; name="uploadfile"; filename="upload_file.txt"] Content-Type:[application/octet-stream]]success Header:map[Content-Length:[130] Content-Type:[text/plain; charset=utf-8] Date:[Fri, 05 Feb 2021 09:47:25 GMT]]}
	log.Printf("%+v", res.Body) // map[Content-Disposition:[form-data; name="uploadfile"; filename="upload_file.txt"] Content-Type:[application/octet-stream]]success
}

func main() {
	// 运行时请先开始example/service/mock_service服务
	uploadRequest()
}
