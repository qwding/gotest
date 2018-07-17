package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd, errCmd := exec.LookPath("cd")
	fmt.Println("cmd ", cmd)
	if errCmd != nil {
		fmt.Println("[ errCmd ]", errCmd)
	}
	commend := exec.Command("cd", "testFolder")
	errCommend := commend.Run()
	if errCommend != nil {
		fmt.Println("[ errCommend ]", errCommend)
	}
}
