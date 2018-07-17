package main

import "fmt"

//两个都是false 提示 map只能跟nil比较。
func main() {
	map1 := map[string]string{"ding": "hit", "yao": "hit"}
	map2 := map[string]string{"ding": "hit", "long": "hit"}
	map3 := map[string]string{"ding": "hit", "yao": "hit"}
	fmt.Println("map1 == map2", map1 == map2)
	fmt.Println("map1 == map3", map1 == map3)
}
