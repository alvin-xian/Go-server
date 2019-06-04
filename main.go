package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type BaseJsonBean struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func NewBaseJsonBean() *BaseJsonBean {
	return &BaseJsonBean{}
}

func loginTask(w http.ResponseWriter, req *http.Request) {
	fmt.Println("loginTask is running...", req.Method)
	//模拟延时
	time.Sleep(time.Second * 2)

	//获取客户端通过GET/POST方式传递的参数
	req.ParseForm()
	param_userName, found1 := req.Form["userName"]
	param_password, found2 := req.Form["password"]

	userName := ""
	password := ""

	if !(found1 && found2) {
		fmt.Fprint(w, "have no userName or password at body")
		if (len(req.Header.Get("userName")) <= 0 || len(req.Header.Get("password")) <= 0) {
			fmt.Fprint(w, "have no userName or password at header")
			return
		} else {
			userName = req.Header.Get("userName")
			password = req.Header.Get("password")
		}
	} else {
		userName = param_userName[0]
		password = param_password[0]
	}

	result := NewBaseJsonBean()

	s := "userName:" + userName + ",password:" + password
	fmt.Println(s)

	if userName == "alvin" && password == "123456" {
		result.Code = 100
		result.Data = "test"
		result.Message = "login ok!"
	} else {
		result.Code = 101
		result.Message = "you are super red!!!"
	}

	//向客户端返回JSON数据
	bytes, _ := json.Marshal(result)
	fmt.Fprint(w, string(bytes))
}

func main() {
	fmt.Println("start:")
	http.HandleFunc("/login", loginTask)
	http.ListenAndServe("10.10.18.60:8001", nil)
	select {} //阻塞进程
}
