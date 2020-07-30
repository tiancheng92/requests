package main

import (
	"fmt"
	"log"
	"requests"
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
