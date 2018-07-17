package main

import (
	"fmt"
	"path"
	"path/filepath"
)

func main() {
	pathc := "imageid://XESORIZaXnC7w1TxNsQs7g=="
	fmt.Println("pathc", pathc)
	fmt.Println("clean", path.Clean(pathc))
	// fmt.Println(path.Join(pathc, "d:/testcode"))
	fmt.Println("if abs path", path.IsAbs("d:/testcode"))
	fmt.Println("dir", path.Dir(pathc))
	fmt.Println("base", path.Base(pathc))
	mypath := "d:/testcode/path"
	fmt.Println("basemy", path.Base(mypath))
	fmt.Println("dirmy", path.Dir(mypath))

	fmt.Println("filepath base:", filepath.Base(pathc))

	/*
		pathc /Users/qwding/gopath/src/go-testcode/path
		clean /Users/qwding/gopath/src/go-testcode/path
		if abs path false
		dir /Users/qwding/gopath/src/go-testcode
		base path
		basemy path
		dirmy d:/testcode
	*/
}
