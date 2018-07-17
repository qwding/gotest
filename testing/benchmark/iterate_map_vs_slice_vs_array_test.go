package testings

import (
	"strconv"
	"testing"
)

var length = 1000
var maps map[string]string
var slices []string
var arrays [1000]string

func init() {
	maps = make(map[string]string, length)
	slices = make([]string, length)
	for i := 0; i < length; i++ {
		maps[strconv.Itoa(i)] = "abc"
		slices[i] = "abc"
		arrays[i] = "abc"
	}

}

func BenchmarkIterateMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, _ = range maps {
		}
	}
}

func BenchmarkIterateSlices(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, _ = range slices {
		}
	}
}

func BenchmarkIterateArrays(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, _ = range arrays {
		}
	}
}

func BenchmarkIterateMapRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, k := range maps {
			_ = k
		}
	}
}

func BenchmarkIterateSlicesRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, k := range slices {
			_ = k
		}
	}
}

func BenchmarkIterateArraysRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, k := range arrays {
			_ = k
		}
	}
}

func BenchmarkIterateSlicesFor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < length; j++ {
			_ = slices[j]
		}
	}
}

func BenchmarkIterateArraysFor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < length; j++ {
			_ = arrays[j]
		}
	}
}
