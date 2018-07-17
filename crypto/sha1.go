package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

//这种方式返回值是有size的[size]byte，这样强制转换不好使了。
/*
func main() {
	name := []byte("dingding")
	res := sha1.Sum(name)
	fmt.Printf("before is %s , after is % x.\n", string(name), res)
	hex_res := hex.EncodeToString(res)
	fmt.Printf("encoding res is %s \n", hex_res) //报错
}*/

func main() {
	start := "dingding"
	h := sha1.New()
	h.Write([]byte(start))
	after := h.Sum(nil)

	fmt.Printf("before is: %s, after is: %s.\n", start, after)

	fmt.Printf("hex is:%s\n", hex.EncodeToString(after))

}
