package main

import "fmt"

type TestStruct struct {
	Name  string
	Count int
}

func main() {
	var maps map[string]TestStruct = make(map[string]TestStruct, 0)
	maps["test1"] = TestStruct{Name: "test1", Count: 3}
	maps["test2"] = TestStruct{Name: "test2", Count: 5}
	if maps["test3"].Name == "" {
		fmt.Println("test3 is nil")

	}
	fmt.Printf("test3 is : %#v\n", maps)
}
