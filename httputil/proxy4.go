package main

import (
	"flag"
	"fmt"
	"github.com/kabukky/httpscerts"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	// "strings"
)

var class string
var port string

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("come into this ")
	o := new(http.Request)

	*o = *r

	// log.Printf("r %#v\n", r)

	destUrlStr := r.Host

	fmt.Println("dest url string", destUrlStr)

	var targetStr string = class + "://" + destUrlStr

	fmt.Println("dest url string", targetStr)

	destUrl, errParse := url.Parse(targetStr)
	if errParse != nil {
		log.Println("[ Error errParse ]", errParse)
	}

	// log.Printf("destUrl %#v\n", destUrl)

	o.Host = destUrl.Host
	o.URL.Scheme = destUrl.Scheme
	o.URL.Host = destUrl.Host
	o.URL.Path = r.URL.Path

	// log.Printf("o.url.string %#v\n", o)

	o.Proto = "HTTP/1.1"
	o.ProtoMajor = 1
	o.ProtoMinor = 1
	o.Close = false

	fmt.Println("o.url.string", o.URL.String())

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
	flag.Parse()
	// ExistConf()
	arglist := flag.Args()

	// Check if the cert files are available.
	err := httpscerts.Check("cert.pem", "key.pem")
	// If they are not available, generate new ones.
	if err != nil {
		err = httpscerts.Generate("cert.pem", "key.pem", "119.81.60.235"+port)
		if err != nil {
			log.Fatal("Error: Couldn't create https certs.")
		}
	}
	if len(arglist) <= 0 {
		fmt.Println("Please give http or https")
		return
	}
	class = arglist[0]
	if class == "https" {
		port = ":443"
	} else if arglist[0] == "http" {
		port = ":80"
	} else {
		fmt.Println("Please give http or https")
		return
	}

	http.HandleFunc("/", handler)

	log.Println("Start serving on port ", port)

	if class == "http" {
		http.ListenAndServe(port, nil)

	} else if class == "https" {
		http.ListenAndServeTLS(port, "cert.pem", "key.pem", nil)

	} else {
		fmt.Println("unbeliveable place")
	}
	fmt.Println("server end")
	os.Exit(0)
}
