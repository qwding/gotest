package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var Num = 10

func main() {
	for i := 0; i < Num; i++ {
		go GenAlert(i)
	}
	select {}
}

type Req struct {
	Labels      map[string]string `json:"labels"`
	Annotations struct {
		Description string `json:"description"`
	} `json:"annotations"`
	Settings struct {
		Overlook string `json:"overlook"`
		Interval string `json:"interval"`
		Timeout  string `json:"timeout"`
	} `json:"settings"`
}

func GenAlert(i int) {
	alert := Req{}
	labels := make(map[string]string)
	labels[string(Krand(5, KC_RAND_KIND_LOWER))] = string(Krand(6, KC_RAND_KIND_ALL))
	labels[string(Krand(5, KC_RAND_KIND_LOWER))] = string(Krand(6, KC_RAND_KIND_ALL))
	labels[string(Krand(5, KC_RAND_KIND_LOWER))] = string(Krand(6, KC_RAND_KIND_ALL))
	labels["user"] = "qwding"

	alert.Labels = labels
	alert.Annotations.Description = "this is num " + strconv.Itoa(i) + "alert."
	alert.Settings.Timeout = "30m"
	alert.Settings.Interval = "60m"

	b, err := json.Marshal([]Req{alert})
	if err != nil {
		fmt.Println("Internal error:", err)
	}

	// fmt.Println("debug body:", string(b))

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// var timer = time.NewTimer(r.Intn(60) * time.Minute)
	timer := time.NewTimer(3 * time.Second)
	j := 0

	for {

		select {
		case <-timer.C:
			client := http.Client{}
			body := bytes.NewBuffer([]byte(b))
			req, err := http.NewRequest("POST", "http://ac.v2.yoo.yunpro.cn/api/alerts", body)
			if err != nil {
				fmt.Println("error,", err)
			}
			req.Header.Add("user", "qwding")
			req.Header.Add("token", "21fa45df-1520-4120-87ef-a071db30fda3")

			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("error:", err)
			}

			br, _ := ioutil.ReadAll(resp.Body)
			defer resp.Body.Close()

			fmt.Println("debug, ", i, j, " req:", string(b), " reap:", string(br))
			timer.Reset(time.Duration(r.Intn(60)) * time.Minute)

		}
		j++
	}
}

// b, err := json.Marshal(result)
// if err != nil {
// 	fmt.Println("Internal error:", err)
// }
// body := bytes.NewBuffer([]byte(b))

// req, err := http.NewRequest("POST", ADDR+"/hosts/create", body)
// if err != nil {
// 	fmt.Println("Internal error:", err)
// }

const (
	KC_RAND_KIND_NUM   = 0 // 纯数字
	KC_RAND_KIND_LOWER = 1 // 小写字母
	KC_RAND_KIND_UPPER = 2 // 大写字母
	KC_RAND_KIND_ALL   = 3 // 数字、大小写字母
)

// 随机字符串
func Krand(size int, kind int) []byte {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if is_all { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return result
}

/*`
[
   {
         "labels":{
             "__name__": "node_memory_Active",
             "alertname": "MemoryActiveOutB",
             "instance": "10.10.12.18:9100",
             "job": "node",
             "monitor": "codelab-monitor",
             "team":"cloud",
             "user":"testcase"
         },
         "annotations": {
             "description": "10.10.12.18:9100 of node memory active out "
         },
         "alerting": true,
         "settings":{
               "overlook":"15s",
               "interval":"30m",
               "timeout":"20m"
         }

     }
 ]
`*/
