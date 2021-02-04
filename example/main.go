package main

import (
	"log"
	"requests"
)

func uploadFile() {
	r := requests.New()
	res, err := r.SetUrl("http://127.0.0.1:9090/upload/").UploadFile("uploadfile", "./file.txt").Post()
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v", res)
}

func getData() {
	r := requests.New()
	res, err := r.SetUrl("http://www.baidu.com/").Get()
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v", res.Body)
}

func main() {
	// uploadFile()
	// getData()
	// ...
}
