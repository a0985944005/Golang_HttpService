package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// 設置路由跟處理方法
	http.HandleFunc("/", HelloWordHandler)
	// 設置監聽的port
	log.Fatal(http.ListenAndServe(":5407", nil))
	//127.0.0.1:5407   localhost:5407
}

func HelloWordHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello Word!!\n")
}
