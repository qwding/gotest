package main

import (
	// "fmt"
	"github.com/shurcooL/markdownfmt/markdown"
	"log"
	"os"
)

func main() {
	input := []byte(`Title
==

Subtitle
---

How about ` + "`this`" + ` and other stuff like *italic*, **bold** and ***super    extra***.
`)

	output, err := markdown.Process("", input, nil)
	if err != nil {
		log.Fatalln(err)
	}

	os.Stdout.Write(output)

	// fmt.Println("res is ", string(output))

}
