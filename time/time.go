package main

import (
	"fmt"
	"time"
)

func main() {
	//
	// local, err := time.LoadLocation("UTC")
	// if err != nil {
	// 	fmt.Println("err ", err)
	// }

	// tmp := time.Date(0, time.January, 0, 78, 0, 0, 0, local)
	// fmt.Println("tmp", tmp)

	// timenow := time.Now()
	// timestr := timenow.Format("2006-01-02")
	// fmt.Println("time ", timestr)

	// 时间戳转换
	// 1473658965
	// t := time.Unix(1473658965, 0)
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().UnixNano())

}
