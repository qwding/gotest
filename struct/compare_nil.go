package main 

import (
"fmt"
)

type SS struct{
	A string
	B int
	C CC
}

type CC struct {
	A string
}
func main(){
	ss := SS{}
	
	if (SS{}) == ss {
		fmt.Println("ss 是空")
	}else{
		fmt.Println("第一次不是空")
	}
	
	ss.C.A = "xxxx"
	if (SS{}) == ss {
                fmt.Println("ss是空")
        }else{
		fmt.Println("第二次不是空")
	}
	s2 := SS{}
	s2.C.A = "xxxx"

	if ss == s2 {
		fmt.Println("ss == s2")
	}else{
		fmt.Println("ss!=s2")
	}
}
