package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
	// "os"
)

type RedisUser struct {
	Username    string
	Password    string
	Email       string
	Permissions map[string]string
	Duration    time.Duration
}

/*func main() {
	fmt.Println("come int")
	jsonstr := `{"username":"edcapding","password":"9e5c73","email":"edcapding@163.com","permissions":{"edcapding":"admin"}}`
	var maps map[string]interface{}
	err := json.Unmarshal([]byte(jsonstr), &maps)
	if err != nil {
		fmt.Println("err ", err)
	}
	// fmt.Println("maps", maps["qwer"])
	fmt.Printf("maps %#v\n", maps)

	v := maps["container0"].(map[string]interface{})
	z := v.(map[string]string)
	fmt.Printf("z is %#v\n", z)
}
*/

func main() {
	fmt.Println("come int")
	err2 := errors.New("{\"username\":\"nkwanglei\",\"duration\":2000000000}")

	// err2 := errors.New(`{"username":"nkwanglei","password":"ad1ead65044d6471f9387532b99d0ee5c61eb990"}`)
	// jsonstr := ` {"username":"nkwanglei","password":"ad1ead65044d6471f9387532b99d0ee5c61eb990","permissions":{"nkwanglei":{"repo":"nkwanglei","access":"admin"}}}`
	// json2 := `{"username":"nkwanglei","password":"ad1ead65044d6471f9387532b99d0ee5c61eb990"}`
	target := &RedisUser{}
	err := json.Unmarshal([]byte(err2.Error()), &target)
	if err != nil {
		fmt.Println("err ", err)
	}
	// fmt.Println("target", target["qwer"])
	fmt.Printf("target %#v\n", target)
	tmp := time.Second * 1
	fmt.Println("tmp:", tmp.String())
	fmt.Println("duration:", target.Duration.String())
}
