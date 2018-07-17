package main

import (
	"fmt"
	// "strings"
	// "os"
	"regexp"
)

func main() {
	/*	for {
		var str string
		fmt.Scanln(&str)
		match, err := regexp.MatchString("^[a-zA-Z]{1}[0-9a-zA-Z_]{1,19}$", str)
		if err != nil {
			fmt.Println("err", err)
		}
		fmt.Println("res ", match)
	}*/
	str := `http://www.ruanyifeng.com/blog/developer/">`
	urlMatch := regexp.MustCompile(`(http://)?[0-9a-zA-Z_\.#-]*[^/]`)
	res := urlMatch.FindAllString(str, 1)
	fmt.Println("res is ", res)
}
