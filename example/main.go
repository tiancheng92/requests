package main

import (
	"fmt"
	"log"
	"requests"
	"time"
)

func ChainCall() {
	var r requests.Request
	res, err := r.SetUrl("http://baidu.com").Get()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%#v\n", res)
}

func StructCall() {
	res, err := requests.Request{
		Method: "GET",
		URL:    "http://baidu.com",
		Data:   "",
		Header: nil,
	}.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%#v\n", res)
}

func StructureCall() {
	var r requests.Request
	r.SetMethod("GET")
	r.SetUrl("http://baidu.com")
	// ...
	res, err := r.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%#v\n", res)
}

type UserInfo struct {
	Count    int64       `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		CellPhone    string        `json:"cell_phone"`
		DateJoined   string        `json:"date_joined"`
		Email        string        `json:"email"`
		ExtendUserID string        `json:"extend_user_id"`
		GroupList    []interface{} `json:"group_list"`
		Groups       []interface{} `json:"groups"`
		ID           int64         `json:"id"`
		IsActive     bool          `json:"is_active"`
		IsStaff      bool          `json:"is_staff"`
		IsSuperuser  bool          `json:"is_superuser"`
		LastLogin    string        `json:"last_login"`
		LastName     string        `json:"last_name"`
		Leader       struct {
			LeaderID   int64  `json:"leader_id"`
			LeaderName string `json:"leader_name"`
		} `json:"leader"`
		Remark interface{} `json:"remark"`
		Team   struct {
			Alias     string `json:"alias"`
			ID        int64  `json:"id"`
			Name      string `json:"name"`
			Structure int64  `json:"structure"`
		} `json:"team"`
		UpdatedAt   int64  `json:"updated_at"`
		User        int64  `json:"user"`
		Username    string `json:"username"`
		Userprofile int64  `json:"userprofile"`
	} `json:"results"`
}

func BindToStruct() {
	var u UserInfo
	var r requests.Request
	res, err := r.SetUrl("").Get()
	if err != nil {
		log.Fatal(err)
		return
	}
	err = res.JsonBind(&u)
	if err != nil {
		return
	}
	fmt.Printf("%s", u.Results[0].ExtendUserID)
}

func SetTimeOut() {
	r := requests.New()
	res, err := r.SetUrl("http://sso.hupu.io").SetTimeOut(time.Second).Get()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%#v\n", res)
}

func main() {
	// ChainCall()
	// StructCall()
	// StructureCall()
	// BindToStruct()
	SetTimeOut()
}
