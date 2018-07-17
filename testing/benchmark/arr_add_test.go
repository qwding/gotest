package testings

import (
	"testing"
	// "fmt"
)

var max = 1000000

func BenchmarkArrAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := make([]string, 0)
		for i := 0; i < max; i++ {
			arr = append(arr, "abc")
		}
	}
}

func BenchmarkArrPoint(b *testing.B) {
	for k := 0; k < b.N; k++ {
		arr := make([]string, 1)
		length := 1
		for i := 0; i < max; i++ {
			if i >= length {
				tmp := make([]string, 2*length)
				length = 2 * length
				for j := 0; j < i; j++ {
					tmp[j] = arr[j]
				}
				arr = tmp
			}
			// fmt.Println("i:",i,"len(arr)",len(arr))
			arr[i] = "abc"
		}
	}
}
