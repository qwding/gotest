package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	//127.0.0.1:8008
	var url string
	for {
		fmt.Println("")
		fmt.Println("please give your url")
		fmt.Scanln(&url)
		cmd := exec.Command("curl", url)
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			fmt.Println("[ Error ]", err)
		}

	}
}
