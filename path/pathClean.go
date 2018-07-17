package main

import (
	"fmt"
	"path"
)

func main() {
	paths := []string{
		"doc/",
		"bbb/xyz",
		"..",
		"doc/..",
	}

	for _, p := range paths {
		fmt.Printf("Clean(%q) = %q\n", p, path.Clean(p))
	}

}
