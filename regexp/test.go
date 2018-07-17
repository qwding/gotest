package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := `2016/12/19 19:41:28 [error] 13871#0: *2635430 [lua] access_by_lua(jsprx.conf:225):142: see code200nilend while sending to client, client: 10.11.1.9, server: _, request: "GET /?d=c095015ccb79eed15928d06d545cafcd&i=36.47.218.195&m=68F7281E1634&redirect=http://www.cnki.com.cn/cnki/js/hotword/hotword.js?fdsaf HTTP/1.1", host: "cachetest.changjingyi.cn", referrer: "http://wiki.cnki.com.cn/HotWord/168471.htm"`
	// text := `2016/12/19 19:41:27 [info] 13869#0: *2636619 [lua] access_by_lua(jsprx.conf:225):158: result: target. while sending to client, client: 10.11.1.1, server: _, request: "GET /?d=c095015ccb79eed15928d06d545cafcd&i=1.80.24.90&m=00012E3E6B90&redirect=http://tui.kugou.com/static/js/music_spread_v4.js?201601041508 HTTP/1.1", host: "cachetest.changjingyi.cn", referrer: "http://tui.kugou.com/html/music_spread_v4.html?keyword=&hash=00000000000000000000000000000000&ver=8123"`

	// reg := regexp.MustCompile(`.*see code.*(client: ).*(GET ).*(host: ).*`)
	// fmt.Printf("%q\n", reg.ReplaceAllString(text, "$1$2$3"))
	// fmt.Println(reg.MatchString(text))
	// "Go 世界！123 Hello."

	match, _ := regexp.MatchString("see code", text)

	if !match {
		fmt.Println("end !")
		return
	}

	// reg2 := regexp.MustCompile(`client: ([0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3})`)
	// fmt.Printf("%q", reg2.FindAll([]byte(text), -1))
	// ["Hello" "World"]

	// reg3 := regexp.MustCompile(`client: ([0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3})`)

	// reg3 := regexp.MustCompile(`GET (\/\?d=.*) HTTP`)

	// reg3 := regexp.MustCompile(`&i=([0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3})`)

	// reg3 := regexp.MustCompile(`&redirect=((http|ftp|https)://)(([a-zA-Z0-9\._-]+\.[a-zA-Z]{2,6})|([0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}))(:[0-9]{1,4})*(/[a-zA-Z0-9\&%_\./-~-]*)?`)
	reg3 := regexp.MustCompile(`host: "(([a-zA-Z0-9\._-]+\.[a-zA-Z]{2,6}))(:[0-9]{1,4})*`)

	// ret := reg3.FindString(text)[10:]

	// (   ([a-zA-Z0-9\._-]+\.[a-zA-Z]{2,6})    |)(:[0-9]{1,4})

	// reg3 := regexp.MustCompile(`(([a-zA-Z0-9\._-]+\.[a-zA-Z]{2,6})|([0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}))(:[0-9]{1,4})`)
	// ret := reg3.FindStringSubmatch(text)
	// reg3 := regexp.MustCompile(`http://([a-zA-Z0-9\._-]+\.[a-zA-Z]{2,6}) `)

	ret := reg3.FindAllString(text, -1)

	// ret := reg3.FindString(text)

	fmt.Println("length:", len(ret), "！！！！！！result:", ret)

	// fmt.Println("result:", ret[0], "第二个:", ret[1])

}
