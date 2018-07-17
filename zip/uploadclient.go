package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func postFile(filename string, targetUrl string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	//关键的一步操作
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}
	//打开文件句柄操作
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	client := &http.Client{}
	request, errRequest := http.NewRequest("POST", targetUrl, bodyBuf)
	if errRequest != nil {
		fmt.Println("[errRequest]", errRequest)
	}
	request.Header.Set("Authority", "Basic ZGluZ2Rpbmc6cXdlcg==")
	request.Header.Set("Content-Type", contentType)
	// resp, err := http.Post(targetUrl, contentType, bodyBuf)

	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))

	return nil
}

// sample usage
func main() {
	target_url := "http://localhost:8080"
	filename := "./123.go"
	postFile(filename, target_url)
}
