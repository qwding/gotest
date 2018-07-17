package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

func main() {
	apiUrl := "https://api.com"
	resource := "/user/"
	data := url.Values{}
	data.Set("name", "foo")
	data.Add("surname", "bar")

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := fmt.Sprintf("%v", u) // "https://api.com/user/"

	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlStr, bytes.NewBufferString(data.Encode())) // <-- URL-encoded payload
	r.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, _ := client.Do(r)
	fmt.Println(resp.Status)
}
