package main

import (
	"fmt"
	"path"
	"path/filepath"
)

func main() {
	fmt.Println(path.Join("bbb", "teest2"))
	fmt.Println(filepath.Join("bbb", "test2"))
}
