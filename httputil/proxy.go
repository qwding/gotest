package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var curUrl string = ""

var urlMap map[string]string = map[string]string{
	"baidu":     "https://www.baidu.com",
	"tenxcloud": "https://www.tenxcloud.com"}

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}

func handler(w http.ResponseWriter, r *http.Request) {
	o := new(http.Request)

	// fmt.Printf("r.request url %#v\n", r.URL)
	// fmt.Printf("r.url %#v\n", r.URL.String())
	fmt.Printf("r %#v\n", r)

	*o = *r
	// fmt.Println("o.url.string before", o.URL.String())

	// fmt.Println("o.url.path", o.URL.Path)

	requestPath := strings.Split(o.URL.String(), "/")

	// fmt.Printf("request array %#v\n", requestPath)

	if len(requestPath) > 1 && urlMap[requestPath[1]] != "" {
		curUrl = urlMap[requestPath[1]]
		requestPath = append(requestPath[:0], requestPath[2:]...)
		o.URL.Path = strings.Join(requestPath, "/")
	}

	targetURL, errParseUrl := url.Parse(curUrl)
	if errParseUrl != nil {
		fmt.Println("[ Error Parse Url ]", errParseUrl)
	}

	// fmt.Println("targetUrl", targetURL)
	// fmt.Println("o.url.string", o.URL.String())

	o.Host = targetURL.Host
	o.URL.Scheme = targetURL.Scheme
	o.URL.Host = targetURL.Host
	o.URL.Path = singleJoiningSlash(targetURL.Path, o.URL.Path)

	// fmt.Println("o.url.RawQuery", o.URL.RawQuery)

	if q := o.URL.RawQuery; q != "" {
		o.URL.RawQuery = o.URL.Path + "?" + q
	} else {
		o.URL.RawQuery = o.URL.Path
	}

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

	log.Println("Start serving on port 1234")

	http.ListenAndServe(":1234", nil)

	os.Exit(0)
}
