package main

import (
	"fmt"
	// "log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./doc")))
	// http.HandleFunc("/", sayHello)
	fmt.Println("start the server1")
	http.ListenAndServe(":8080", nil)

	fmt.Println("start the server2")
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello ding")
}
