package main

import (
	"fmt"
)

func main() {
	orgin := []string{"xyz", "opq", "qwer", "xyz", "qerw", "qwer"}
	ret := TrimSame(orgin)
	fmt.Println("result is :", ret)

}

func TrimSame(urls []string) []string {
	lengthurl := len(urls)
	for i := 0; i < lengthurl; i++ {
		for j := i + 1; j < lengthurl; j++ {
			if urls[i] == urls[j] {
				urls = append(urls[:j], urls[j+1:]...)
				lengthurl--
				continue
			}
		}
	}
	return urls
}
