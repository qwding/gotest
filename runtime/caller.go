package main

import (
	"fmt"
	"runtime"
)

func main() {
	for skip := 0; ; skip++ {
		pc, file, line, ok := runtime.Caller(skip)
		if !ok {
			break
		}
		fmt.Printf("skip = %v, pc = %v, file = %v, line = %v\n", skip, pc, file, line)
	}
	// Output:
	// skip = 0, pc = 4198453, file = caller.go, line = 10
	// skip = 1, pc = 4280066, file = $(GOROOT)/src/pkg/runtime/proc.c, line = 220
	// skip = 2, pc = 4289712, file = $(GOROOT)/src/pkg/runtime/proc.c, line = 1394
}
