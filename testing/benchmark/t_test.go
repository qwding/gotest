package testings

import (
"fmt"
"testing"
"time"
)

func Test_A(t *testing.T){
	t.Parallel()
	fmt.Println("testing a begin.")
	time.Sleep(time.Second*10)
	fmt.Println("testing a end.")
}

func Test_B(t *testing.T){
	fmt.Println("testing b begin.")
	time.Sleep(time.Second*7)
	fmt.Println("testing b end.")
}

func Test_C(t *testing.T){
	t.Parallel()
	fmt.Println("testing c begin.")
	time.Sleep(time.Second*5)
	fmt.Println("testing c end.")
}