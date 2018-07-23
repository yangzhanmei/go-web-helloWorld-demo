package main

import (
	"fmt"
	"net/http"
	"log"
)

func getTodoList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		w.Header().Set("Access-Control-Allow-Origin", "*");
		w.Write([]byte("Hello World"));

	} else {
		fmt.Println(w);
	}
}


func main() {
	http.HandleFunc("/", getTodoList) //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}