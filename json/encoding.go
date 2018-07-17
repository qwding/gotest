package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	fmt.Println("come int")
	jsonstr := `{"qwer":1234,"asdf":"zxc"}`
	wr := json.NewEncoder(os.Stdout)
	wr.Encode(jsonstr)
	fmt.Println("over")

	//decode
	var maps map[string]string = make(map[string]string)
	input := os.Stdin
	read := json.NewDecoder(input)
	read.Decode(maps)

	fmt.Println(maps)
}
