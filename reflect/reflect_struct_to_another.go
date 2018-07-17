package main

import (
	"encoding/json"
	"fmt"
)

type Point struct {
	X int
	Y int
}

type Point2 struct {
	X int32
	Y int64
}

type A struct {
	AA string
	AB int
	AC map[string]*Point
}

type B struct {
	AA string
	AB int
	AC map[string]Point2
}

func main() {
	a := A{AA: "a", AB: 123}
	a.AC = map[string]*Point{"aaa": &Point{X: 2, Y: 4}}
	res, err := json.Marshal(a)
	if err != nil {
		println("marshal err:", err)
	}

	b := B{}
	err = json.Unmarshal(res, &b)
	if err != nil {
		println("Unmarshal error:", err)
	}

	fmt.Printf("b %#v\n", b)
}
