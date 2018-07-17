package main

import (
	"fmt"
	"io"
	// "io/ioutil"
	"net/http"
)

const (
	// giturl = "https://git.yunpro.cn"
	giturl = "http://10.50.1.56:10080"
)

var clientid = "1b44e5046e51199065f89f0ad2a59385978cd025bedb1353c9148f7496907804"
var sercetid = "15b139748d066635620491ebfcd0349b5a4b5457ba9d7728e2e3105621b227cc"

func handler(w http.ResponseWriter, req *http.Request) {
	// fmt.Println("StatusTemporaryRedirect:", http.StatusTemporaryRedirect)
	http.Redirect(w, req, giturl+"/oauth/authorize?response_type=code&client_id="+clientid+"&redirect_uri=http%3A%2F%2Flocalhost:8888", http.StatusTemporaryRedirect)
}

// http%3A%2F%2Fwww.jianshu.com%2Fusers%2Fauth%2Fweibo%2Fcallback&response_type=code&state=34f6bea578f01965a386d4699462374893858701a5eb91d6
// http%3A%2F%2Flocalhost:8008%2Fcallback
// http%3A%2F%2Fwww.baidu.com

func test(w http.ResponseWriter, req *http.Request) {
	fmt.Println("in test,callback")
	fmt.Println("response", w)
	fmt.Println("request:", req)
	io.WriteString(w, "1233")
}

//访问 localhost:8008/abc，重定向到gitlab登录，并且进行授权，在重定向返回到 /callback 页面

func main() {
	port := ":8888"
	http.HandleFunc("/", test)
	http.HandleFunc("/abc", handler)
	fmt.Println("Start server at ", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error ", err)
	}
	fmt.Println("End Server")
}
