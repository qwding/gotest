package main

import (
	"fmt"
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(401)
	io.WriteString(w, "<h1 align=\"center\">你大爷！推广的别TM打电话给我</h1>")
}
func main() {
	port := ":8008"
	http.HandleFunc("/", handler)
	fmt.Println("Start server at ", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error ", err)
	}
	fmt.Println("End Server")
}
