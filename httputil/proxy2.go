package main

import (
	// "fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var urlMap map[string]string = map[string]string{
	"www.baidu.com":  "https://www.tenxcloud.com",
	"www.google.com": "https://www.tenxcloud.com",
	"www.beego.me":   "https://www.tenxcloud.com"}

var port string = ":80"

func handler(w http.ResponseWriter, r *http.Request) {
	o := new(http.Request)

	*o = *r

	// fmt.Printf("r %#v\n", r)

	var destUrlStr string
	if port == ":80" {
		destUrlStr = r.Host
	} else {
		destUrlStr = strings.Replace(r.Host, port, "", 1)
	}

	if urlMap[destUrlStr] == "" {
		log.Println("[ Error ] Undefied the URL in program")
		return
	}
	destUrl, errParse := url.Parse(urlMap[destUrlStr])
	if errParse != nil {
		log.Println("[ Error ]", errParse)
	}

	// fmt.Printf("destUrl %#v\n", destUrl)

	o.Host = destUrl.Host
	o.URL.Scheme = destUrl.Scheme
	o.URL.Host = destUrl.Host
	o.URL.Path = r.URL.Path

	// fmt.Printf("o.url.string %#v\n", o)

	o.Proto = "HTTP/1.1"
	o.ProtoMajor = 1
	o.ProtoMinor = 1
	o.Close = false

	// fmt.Println("o.url.string after", o.URL.String())

	transport := http.DefaultTransport

	res, err := transport.RoundTrip(o)

	if err != nil {
		log.Printf("http: proxy error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	hdr := w.Header()

	for k, vv := range res.Header {
		for _, v := range vv {
			hdr.Add(k, v)
		}
	}

	for _, c := range res.Cookies() {
		w.Header().Add("Set-Cookie", c.Raw)
	}

	w.WriteHeader(res.StatusCode)

	if res.Body != nil {
		io.Copy(w, res.Body)
	}
}

func main() {

	http.HandleFunc("/", handler)

	log.Println("Start serving on port ", port)

	http.ListenAndServe(port, nil)

	os.Exit(0)
}
