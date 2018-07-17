package main

import (
	"fmt"
	"os"
)

func main() {
	str := "opq/rst/uvw"
	err := MakeDir(str)
	if err != nil {
		fmt.Println(err)
	}
}

func MakeDir(name string) error {
	err := os.Mkdir(name, 0666)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
