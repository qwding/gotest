package main

import "fmt"

type Test struct {
	Field int
}

//如果不带星，会报错！！！ .\map.go:12: cannot assign to maps["first"].Field
func main() {
	var maps = map[string]*Test{"first": &Test{Field: 2}, "second": &Test{Field: 5}}
	maps["first"].Field = maps["first"].Field + 1
	fmt.Println("res:", maps["first"].Field)
}
