package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

func main() {
	dec := json.NewDecoder(os.Stdin)
	enc := json.NewEncoder(os.Stdout)
	fmt.Println("dec type", reflect.TypeOf(dec))
	fmt.Println("enc type", reflect.TypeOf(enc))
	for {
		var v map[string]interface{}
		err1 := dec.Decode(&v)
		if err1 != nil {

			fmt.Println("1 ", err1)
			return
		}
		for k := range v {
			if k != "Title" {
				v[k] = nil
			}
		}
		// fmt.Println("type of enc encode:", reflect(enc.Encode(&v)))
		if err2 := enc.Encode(&v); err2 != nil {
			fmt.Println("2 ", err2)
		}
	}

}
