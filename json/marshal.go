package main

import (
	"encoding/json"
	"fmt"
	"time"
	// "os"
)

func main() {
	/*maps := make(map[string]string)
	maps["abc"] = "123"
	maps["qwer"] = "345"
	maps["asdf"] = "567"
	res, err := json.Marshal(maps)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("res is ", res)
	fmt.Println("res is ", string(res))*/

	image := Image{Name: "name1", Describe: "des1", ImgPath: "/abc"}
	res, err := json.Marshal(image)
	if err != nil {
		fmt.Println("err", err)
	}

	fmt.Println("res:", string(res))

	fmt.Printf("%#v\n", image)

}

type Image struct {
	ImageId  string         `json:"id"bson:"_id"`
	Tags     map[string]Tag `json:"tags"`
	Name     string         `json:"name"`
	Describe string         `json:"describe"`
	ImgPath  string         `json:"imgpath"`
}

type Tag struct {
	Env        map[string]string
	Ports      []string
	Volumes    []string
	Cmd        []string
	Entrypoint []string
	Layers     int
	Created    time.Time
}
