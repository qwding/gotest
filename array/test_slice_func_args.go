package main

import (
	"fmt"
	"unsafe"
)

// func main() {
// 	result := make([]int, 0, 10)
// 	addr := &result
// 	fmt.Println("result:", result, "addr:", unsafe.Pointer(addr))

// 	functions(result)
// 	fmt.Println("result:", result, "addr:", unsafe.Pointer(addr))

// }

// func functions(arr []int) {
// 	for i := 0; i < 10; i++ {
// 		addr := &arr

// 		fmt.Println("arr:", arr, "addr:", unsafe.Pointer(addr))
// 		arr = append(arr, i)
// 	}
// }

//这段代码发现，在函数里是不能增加元素的，因为是值传递
// func main() {
// 	result := make([]int, 0, 10)
// 	addr := &result
// 	fmt.Println("result:", result, "addr:", unsafe.Pointer(addr))

// 	for i := 0; i < 10; i++ {
// 		functions(result)
// 	}
// 	fmt.Println("result:", result, "addr:", unsafe.Pointer(addr))

// }

// func functions(arr []int) {
// 	addr := &arr
// 	fmt.Println("arr:", arr, "addr:", unsafe.Pointer(addr))
// 	arr = append(arr, 8)
// }

//这说明，切片里的元素的地址是不变的，所以切片传指针是因为切片元素的地址传进来了，而切片本身地址并没传进来
//所以说，切片传递时候是不可加，但是可修改的
func main() {
	result := make([]int, 1, 11)
	addr := &result
	fmt.Println("result:", result, "addr:", unsafe.Pointer(addr))

	for i := 0; i < 10; i++ {
		functions(result)
		result = append(result, i)
	}
	fmt.Println("result:", result, "addr:", unsafe.Pointer(addr))

}

func functions(arr []int) {
	addr := &arr
	if len(arr) > 0 {
		fmt.Println("arr:", arr, "addr:", unsafe.Pointer(addr), "arr 0 addr:", &arr[0])
	} else {
		fmt.Println("arr:", arr, "addr:", unsafe.Pointer(addr))
	}
}

//这说明切片传递时候是不可删的
// func main() {
// 	result := []int{1, 2, 3, 4, 5, 5, 6, 6}
// 	addr := &result
// 	fmt.Println("result:", result, "addr:", unsafe.Pointer(addr))

// 	functions(result)
// 	fmt.Println("result:", result, "addr:", unsafe.Pointer(addr))

// }

// func functions(arr []int) {
// 	l := len(arr)
// 	arr = arr[0 : l/2]
// }
