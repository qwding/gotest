package main

import "fmt"
import "github.com/howeyc/gopass"

func main() {
	fmt.Printf("Password: ")
	pass := gopass.GetPasswdMasked()
	fmt.Println("res", string(pass))
	// Do something with pass
}
