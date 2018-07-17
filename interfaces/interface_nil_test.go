package main

import (
	"fmt"
	"testing"
)

type Cat interface {
	Meow()
}

type Tabby struct{}

func (*Tabby) Meow() { fmt.Println("meow") }

func GetACat() Cat {
	var myTabby *Tabby = nil
	// 天哦，忘记给 myTabby 赋值了
	return myTabby
}

func TestGetACat(t *testing.T) {
	if cat := GetACat(); cat == nil {
		t.Errorf("忘记返回一个真实的 Cat!")
	} else {
		fmt.Println("in else", cat)
	}
}
