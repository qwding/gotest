package main

import (
	"fmt"
	"os/user"
)

func main() {
	u, err := user.Current()
	if err != nil {
		fmt.Println("err:", err)
	}

	fmt.Printf("user %#v\n", u)

}
