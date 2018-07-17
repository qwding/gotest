package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func handler(w http.ResponseWriter, req *http.Request) {
	url := req.URL
	fmt.Println("url:", url)

	io.WriteString(w, "<h1 align=\"center\">你大爷！推广的别TM打电话给我</h1>")
}

func test(w http.ResponseWriter, req *http.Request) {
	fmt.Println("in test,callback")
	fmt.Println("response", w)
	fmt.Println("request:", req)

	body := req.Body
	defer req.Body.Close()

	b, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("body:", string(b))

	io.WriteString(w, "1233")
}

func main() {
	port := ":8008"
	http.HandleFunc("/callback", test)
	http.HandleFunc("/abc", handler)
	fmt.Println("Start server at ", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error ", err)
	}
	fmt.Println("End Server")
}

func post() {
	b, err := json.Marshal(result)
	if err != nil {
		fmt.Println("Internal error:", err)
	}
	body := bytes.NewBuffer([]byte(b))

	req, err := http.NewRequest("POST", ADDR+"/hosts/create", body)
	if err != nil {
		fmt.Println("Internal error:", err)
	}

}
