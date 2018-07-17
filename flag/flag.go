package main

import (
	"flag"
	"fmt"
)

var localStorage = flag.Bool("local", false, "save image to local.")

func main() {
	all := flag.Bool("a", false, "set all")
	flag.Parse()
	args := flag.Args()
	for i, j := range args {
		fmt.Println("args ", i, j)
	}

	if *localStorage {
		fmt.Println("local storage:", *localStorage)
	} else {
		fmt.Println("local not storage:", *localStorage)
	}

	fmt.Println("exist -a", *all)
	fmt.Println("end")
}
