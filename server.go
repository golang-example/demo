package main

import (
	"net/http"
	"fmt"
	"log"
)

func main() {
	//http请求处理
	http.HandleFunc("/test", test)

	fmt.Print("http server start.....")
	//绑定监听地址和端口
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}
