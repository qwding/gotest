package main

import (
	"fmt"
)

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}

	fmt.Println("stu:", stus)

	for _, stu := range stus {
		fmt.Println("iiii:", stu, "m:", stu)
		m[stu.Name] = &stu
	}
	fmt.Println("m:", m, "arr:", stus)
	println(m["zhou"], m["li"], m["li"])

	m["zhou"] = &student{Name: "zhou", Age: 28}

	fmt.Println("m:", m, "arr:", stus)
	fmt.Println("zhou:", m["zhou"])
}

func main() {
	pase_student()
}
