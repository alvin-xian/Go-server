package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type BaseJsonBean struct {
	Code    int       //  `json:"code"`
	Data    string// interface{} //`json:"data"`
	Message string      //`json:"message"`
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

	fmt.Println("Http info:", req.Method,"req.Form:", req.Form, "req.Body:", req.Body, "req.Header:", req.Header, "\n req.PostForm:", req.PostForm)

	userName := ""
	password := ""

	if !(found1 && found2) {
		//查看格式是否是mutipartform, 需要用ParseMultipartForm的方式来届时
		req.ParseMultipartForm(1024)
		fmt.Println("ParseMultipartForm", req.Form,req.MultipartForm.Value);
		param_userName, found1 = req.Form["userName"]
		param_password, found2 = req.Form["password"]
		if !(found1 && found2) {
			//查看数据是否存放在header,一般不建议账号密码存放在header。
			if (len(req.Header.Get("userName")) <= 0 || len(req.Header.Get("password")) <= 0) {
				fmt.Fprint(w, "have no userName or password", "\n")
				return
			} else {
				userName = req.Header.Get("userName")
				password = req.Header.Get("password")
				fmt.Fprint(w, "get info on header", "\n")
			}
		}else{
			//数据放在body上上传
			userName = param_userName[0]
			password = param_password[0]
			fmt.Fprint(w, "get info on body with mutipart/form:", req.Form, "\n")
		}
	} else {
		if(req.Method == "GET"){
			fmt.Fprint(w, "get info on url:", req.URL, "\n")
		}else {
			fmt.Fprint(w, "get info on body:", req.Form, "\n")
		}
		//数据放在body上上传
		userName = param_userName[0]
		password = param_password[0]
	}

	result := NewBaseJsonBean()

	s := "userName:" + userName + ",password:" + password
	fmt.Println("login info:",s, "\n")

	if userName == "alvin" && password == "123456" {
		result.Code = 100
		result.Data = "alvin's server"
		result.Message = "login sucessed!"
	} else {
		result.Code = 101
		result.Data = "alvin's server"
		result.Message = "login failed!"
	}

	//向客户端返回JSON数据
	bytes, _ := json.Marshal(result)
	fmt.Fprint(w, string(bytes), "\n")
}

func main() {
	go func() {
		for i:=0;i<10;i++ {
		fmt.Println("test go func", i)
		}
	}()
	fmt.Println("start:")
	http.HandleFunc("/login", loginTask)
	http.ListenAndServe("10.10.18.60:8001", nil)
	go func() {
		fmt.Println("end.")
	}()
	//select {} //阻塞进程
}
