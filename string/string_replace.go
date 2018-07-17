package main

import (
	"fmt"
	"strings"
)

/*func main() {
	r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
	fmt.Println(r.Replace("This is <b>HTML</b>!"))
}
*/
func main() {
	str := "opq_df_nnn_fdsa"
	res := strings.Replace(str, "_", "/", 1)
	fmt.Println("before:", str, "after:", res)
}
