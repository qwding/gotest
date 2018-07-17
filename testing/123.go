package testings

import (
	// "fmt"
	"testing"
)

func Benchmark_ArrLen(b *testing.B) {
	maps := make(map[int]int, 0)
	for i := 0; i < b.N; i++ {
		maps[i] = i
		_ = len(maps)
	}
}

func Benchmark_MapLen(b *testing.B) {
	arr := make([]int, 0)
	for i := 0; i < b.N; i++ {
		arr = append(arr, i)
		_ = len(arr)
	}
}

func Benchmark_Count(b *testing.B) {
	a := 0
	for i := 0; i < b.N; i++ {
		a = a + 1
		_ = a
	}
}