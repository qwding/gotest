package main

import "fmt"
import "unsafe"

func main() {
	a := int(123)
	b := int64(123)
	c := "foo"
	d := struct {
		FieldA float32
		FieldB string
	}{0, "bar"}

	fmt.Printf("a: %T, %d\n", a, unsafe.Sizeof(a))
	fmt.Printf("b: %T, %d\n", b, unsafe.Sizeof(b))
	fmt.Printf("c: %T, %d\n", c, unsafe.Sizeof(c))
	fmt.Printf("d: %T, %d\n", d, unsafe.Sizeof(d))
	fmt.Println("c:", unsafe.Pointer(&c))
	fmt.Println("c addr:", &c)
	fmt.Println("c byte:", []byte(c))
	fmt.Println(0xc8200701c0)
	fmt.Println("c P:", *&c)
}
