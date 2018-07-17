package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "14537Gfi"

	defer func() {
		if rec := recover(); rec != nil {
			fmt.Println("rec is ", rec)
		}
	}()

	end := converMem(str)
	fmt.Println("end", end)
}

func converMem(str string) int {
	if str == "" {
		return 0
	}
	method := "converMem"
	memstr := str[:len(str)-2]
	mem, err := strconv.Atoi(memstr)
	if err != nil {
		fmt.Println(method, err)
	}

	fmt.Println("str is ", str)
	ki := strings.Index(str, "Ki")
	mi := strings.Index(str, "Mi")
	gi := strings.Index(str, "Gi")
	fmt.Println("ki,mi,gi", ki, mi, gi)

	if ki > 0 {
		mem = mem * 1024
	} else if mi > 0 {
		mem = mem
	} else if gi > 0 {
		mem = mem / 1024
	} else {
		panic(fmt.Sprintln(method, "can't find the unit of mem."))
	}
	return mem
}
