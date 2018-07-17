package testings

import (
	"testing"
	// "fmt"
)

func BenchmarkDeferOrigin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() {
			_ = i
		}()
	}
}

func BenchmarkDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() {
			defer func() { _ = i }()
			_ = i
		}()
	}
}
