package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
)

func main() {
	str := "m.youku.com/video/id_XMTczNDAwNjQzNg==.html"

	fmt.Println(GenId(str))
}

func GenId(url string) string {
	tmp := sha256.Sum256([]byte(url))

	url = string(tmp[:])

	id := base64.StdEncoding.EncodeToString([]byte(url))
	fmt.Println("base64:", id)
	id = strings.Replace(id, "+", "-", -1)
	id = strings.Replace(id, "/", "_", -1)
	id = strings.Replace(id, "=", "", -1)
	return id
}
