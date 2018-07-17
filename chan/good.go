package main

import (
	"fmt"
)

// 1.
// func main() {
// 	// 设置流水线
// 	c := gen(2, 3)
// 	out := sq(c)

// 	// 消费输出结果
// 	fmt.Println(<-out) // 4
// 	fmt.Println(<-out) // 9
// }

// 2.
func main() {
	in := gen(2, 3)

	// 启动两个 sq 实例处理
	// channel "in" 发送的数据
	c1 := sq(in)
	c2 := sq(in)

	// merge 函数将 channel c1 和 c2
	// 合并到一起，
	// 这段代码会消费 merge 的结果
	for n := range merge(c1, c2) {
		fmt.Println(n) // 打印 4 9, 或 9 4
	}
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		fmt.Println("chan len:", len(in))
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func gen(nums ...int) <-chan int {

	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}
