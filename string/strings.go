package main

import (
	"fmt"
	// "strconv"
	"strings"
)

// func main() {

// 	//string index
// 	// str1 := "qwertyu"
// 	// str2 := "ert"
// 	// index := strings.Index(str1, str2)
// 	// fmt.Println("index is ", index)

// 	// strconv atoi
// 	str3 := "123"
// 	str4 := "234q"
// 	str3Res, err3 := strconv.Atoi(str3)
// 	str4Res, err4 := strconv.Atoi(str4)
// 	fmt.Println("str3Res", str3Res, err3)
// 	fmt.Println("str4Res", str4Res, err4)

// }

/*func main() {

	//分割的 / 就没有了。第一个是空字符串
	str := "/baidu?o/opq/rst?fdsfhsfhah"
	strArray := strings.Split(str, "/")
	fmt.Printf("res %#v\n", strArray)

}*/

//get length
// func main() {
// 	str := "b5d13319cdf9efb0eb5d42315b88ed2f1da0eb9ac328076d"
// 	fmt.Println("length:", len(str))
// }

func main() {
	str := "fasd,fds,fas,fdsa,fa"
	count := strings.Count(str, ",")
	fmt.Println("count:", count)
}
