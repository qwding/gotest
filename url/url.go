package main

import (
	fm "fmt"
	"net/url"
	rt "runtime"
)

func main() {
	var urlStr string = "http://baidu.com/index.php/?abc=1_羽毛"
	l, err := url.ParseQuery(urlStr)
	fm.Println(l, err) //map[http://baidu.com/index.php/?abc:[1_羽毛]] <nil>
	l2, err2 := url.ParseRequestURI(urlStr)
	fm.Println(l2, err2) //http://baidu.com/index.php/?abc=1_羽毛 <nil>

	l3, err3 := url.Parse(urlStr)
	fm.Println(l3, err3) // http://baidu.com/index.php/?abc=1_羽毛 <nil>
	fm.Println("host", l3)
	fm.Println("path", l3.Path) ///index.php/
	//	fm.Println(l3.RawQuery)         //abc=1_羽毛
	//	fm.Println(l3.Query())          //map[abc:[1_羽毛]]
	//	fm.Println(l3.Query().Encode()) //abc=1_%E7%BE%BD%E6%AF%9B
	l3.Scheme = "fs"
	fm.Println("l3:", l3.String())

	fm.Println(l3.RequestURI())                          ///index.php/?abc=1_羽毛
	fm.Printf("Hello World! version : %s", rt.Version()) //Hello World! version : go1.6.2
}

/*
map[http://baidu.com/index.php/?abc:[1_羽毛]] <nil>
http://baidu.com/index.php/?abc=1_羽毛 <nil>
http://baidu.com/index.php/?abc=1_羽毛 <nil>
/index.php/
abc=1_羽毛
map[abc:[1_羽毛]]
abc=1_%E7%BE%BD%E6%AF%9B
/index.php/?abc=1_羽毛
Hello World! version : go1.6.2
*/
