package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// 設置路由跟處理方法
	http.HandleFunc("/post", PostHandler)
	http.HandleFunc("/get", GetHandler)
	// 設置監聽的port
	log.Fatal(http.ListenAndServe(":5407", nil))
}

type ReqData struct {
	Method  string
	Body    string
	Headers map[string][]string
	Cookie  []*http.Cookie
	Params  map[string][]string
	Url     string
}

func (r ReqData) String() string {
	b, err := json.Marshal(r)
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func PostHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("req method: ", req.Method)
	// req.Method -> http method
	if req.Method == "POST" {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(w, "read request body error", http.StatusInternalServerError)
		}
		reqdata := ReqData{
			Method:  req.Method,
			Body:    string(body),
			Headers: req.Header,
			Params:  req.URL.Query(),
			Cookie:  req.Cookies(),
			Url:     req.URL.String(),
		}

		fmt.Fprint(w, reqdata.String())
		return
	}
	http.Error(w, "Don't use Get / invalid request method", http.StatusMethodNotAllowed)
}

func GetHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("req method: ", req.Method)

	io.WriteString(w, "Hello, It GetHandler!\n")
}
