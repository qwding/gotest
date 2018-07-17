package main

import (
	"fmt"
	"github.com/kabukky/httpscerts"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var urlMap map[string]string = map[string]string{
	"go.googlesource.com": "https://go.googlesource.com",
	"www.google.com":      "https://www.google.com",
	"www.google.cn":       "http://www.google.cn",
	"www.google.com.sg":   "https://www.google.com.sg"}

var port string = ":443"

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("come into this ")
	o := new(http.Request)

	*o = *r

	log.Printf("r %#v\n", r)

	var destUrlStr string
	if port == ":80" || port == ":443" {
		destUrlStr = r.Host
	} else {
		destUrlStr = strings.Replace(r.Host, port, "", 1)
	}

	fmt.Println("desturlstr", destUrlStr)

	if urlMap[destUrlStr] == "" {
		log.Println("[ Error ] Undefied the URL in program")
		return
	}
	destUrl, errParse := url.Parse(urlMap[destUrlStr])
	if errParse != nil {
		log.Println("[ Error errParse ]", errParse)
	}

	log.Printf("destUrl %#v\n", destUrl)

	o.Host = destUrl.Host
	o.URL.Scheme = destUrl.Scheme
	o.URL.Host = destUrl.Host
	o.URL.Path = r.URL.Path

	log.Printf("o.url.string %#v\n", o)

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

	// Check if the cert files are available.
	err := httpscerts.Check("cert.pem", "key.pem")
	// If they are not available, generate new ones.
	if err != nil {
		err = httpscerts.Generate("cert.pem", "key.pem", "127.0.0.1"+port)
		if err != nil {
			log.Fatal("Error: Couldn't create https certs.")
		}
	}

	http.HandleFunc("/", handler)

	log.Println("Start serving on port ", port)

	http.ListenAndServeTLS(port, "cert.pem", "key.pem", nil)

	os.Exit(0)
}
