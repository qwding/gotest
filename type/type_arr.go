package main

import (
	"fmt"
)

type urls []string

func newUrls(start []string) urls {
	return urls(start)
}

func (u urls) inList(url string) bool {
	// for _, val := range []string(u) {
	for _, val := range u {
		if val == url {
			return true
		}
	}
	return false
}
func (u *urls) addList(urls []string) {
	for _, url := range urls {
		u.add(url)
	}
}

/*func (u *urls) add(url string) {
	if !u.inList(url) {
		fmt.Println("urls add,")
		tmp := append([]string(*u), url)
		fmt.Println("urls add,", "after add :", tmp)
		*u = urls(tmp)
		// u = urls(tmp)
	}
}*/

func (u *urls) add(url string) {
	if !u.inList(url) {
		*u = append(*u, url)
		// u = urls(tmp)
	}
}

func main() {
	tmp := newUrls([]string{"abc", "bcd", "123"})
	fmt.Println("if 123 inlist:", tmp.inList("123"))

	tmp.addList([]string{"add3", "add4", "abc"})

	fmt.Println("res:", tmp)
}
