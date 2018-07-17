package main

import (
	"fmt"
	// "html/template"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintln(w, "<html><head><title>TenxCloud Hello World Sample</title></head><body><div align=\"center\"><h1>Hello world! Powered by TenxCloud!</h1><img height=\"100\" src=\"/static/logo.jpg\"></div></body></html>")
		fmt.Println("Hello world, this is TenxCloud!")
		return
	}
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Listen And Server", err.Error())
	}
	fmt.Println("Server Listening on 8080 ")
}
