package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := GetV1Prefix() + "/image/" + "edcapding/beegobase" //nkwanglei/nodejssample private
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Authorization", "Basic ZWRjYXBkaW5nOmRpbmdkaW5n")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	bodyStruct := resp.Body
	defer bodyStruct.Close()
	bodyByte, err := ioutil.ReadAll(bodyStruct)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("body is ", string(bodyByte))

}

func GetV1Prefix() string {
	return "https://index.tenxcloud.com:443/v1"
}
