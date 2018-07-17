package main

import (
	"fmt"
	"reflect"
)

type test struct {
	Abc string
	Xyz string
	Q   int
	S   SS
}
type SS struct {
	A string
}

func main() {
	tt := &test{"abc", "xyz", 2, SS{"fda"}}

	res := GetFields(tt)
	fmt.Printf("res: %#v\n", res)
}

func GetFields(i interface{}) map[string]string {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	res := make(map[string]string)

	for i := 0; i < t.NumField(); i++ {
		fmt.Println("type:", t.Field(i).Type)
		if v.Field(i).Kind() == reflect.String {
			res[t.Field(i).Name] = v.Field(i).String()
		}
	}
	return res
}
