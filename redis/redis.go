package main

import (
	"fmt"
	// 导入redigo扩展包
	"github.com/garyburd/redigo/redis"
	// "strings"
	"reflect"
	// "encoding/json"
)

func main() {
	host := "billing-redis-vq1wu.q1.tenxcloud.net:43507"
	db := "billing"
	password := "billing"
	// tcp连接redis
	rs, err := redis.Dial("tcp", host)
	// 操作完后自动关闭

	// 若连接出错，则打印错误信息，返回
	if err != nil {
		fmt.Println("redis connect error", err)
		return
	}

	err = rs.Send("AUTH", password)
	if err != nil {
		fmt.Println("send password err", err)
	}
	defer rs.Close()
	// 选择db
	rs.Do("SELECT", db)

	for {

		fmt.Println("please give action***************")
		fmt.Println("key list get lrange hash")
		var action string
		fmt.Scanln(&action)
		if action == "key" {
			fmt.Println("please give insert key value")
			var key string
			var value string
			fmt.Scanln(&key)
			fmt.Scanln(&value)
			// 操作redis时调用Do方法，第一个参数传入操作名称（字符串），然后根据不同操作传入key、value、数字等
			// 返回2个参数，第一个为操作标识，成功则为1，失败则为0；第二个为错误信息
			n, err := rs.Do("SETNX", key, value)
			// 若操作失败则返回
			if err != nil {
				fmt.Println("rs.Do err", err)
				return
			}
			// 返回的n的类型是int64的，所以得将1或0转换成为int64类型的再比较
			if n == int64(1) {
				fmt.Println("add successful")
				// 设置过期时间为24小时
				n, _ := rs.Do("EXPIRE", key, 24*3600)
				if n == int64(1) {
					fmt.Println("success")
				}
			} else if n == int64(0) {
				fmt.Println("the key has already existed")
			}
		} else if action == "get" {
			fmt.Println("please give the key you want get")

			var key string
			fmt.Scanln(&key)
			// 调用Do后，还得调用相应的方法才能取得数据
			// 由于之前存的value是string类型，所以用redis.String将数据转换成string类型
			value, err := redis.String(rs.Do("GET", key))
			if err != nil {
				fmt.Println("fail")
			}
			fmt.Println("get the key ", key, "value is ", value)
		} else if action == "list" {
			fmt.Println("please give the list name you want push")
			var key string
			fmt.Scanln(&key)
			fmt.Printf("please give the key %s's value\n", key)
			var value string
			fmt.Scanln(&value)
			n, err := rs.Do("lpush", key, value)
			if err != nil {
				fmt.Println("add list error", err)
			}
			if n == int64(1) {
				fmt.Println("add success")
			} else {
				fmt.Println("add not success,n is", n)
			}
		} else if action == "lrange" {
			fmt.Println("please give the list name you want get")
			var key string
			fmt.Scanln(&key)
			value, err := redis.Values(rs.Do("lrange", key, "0", "1000"))
			if err != nil {
				fmt.Println("get list err", err)
			}
			for _, v := range value {
				fmt.Println("list value is", string(v.([]byte)))
			}
		} else if action == "hash" {
			fmt.Println("give hash key")
			var key string
			fmt.Scanln(&key)
			// arr := []interface{}{"name", "good", "age", 11, "job", "teacher"}

			n, err := rs.Do("hmset", key, "name", "good", "age", 12, "job", "teacher")
			// n, err := rs.Do("hkeys", key)
			// 若操作失败则返回
			if err != nil {
				fmt.Println("rs.Do err", err)
				return
			}

			/*if len(n.([]interface{})) <= 0 {
				fmt.Println("n len is < 0")
			} else {
				fmt.Println("n len is > 0", len(n.([]interface{})))
			}*/

			// 返回的n的类型是int64的，所以得将1或0转换成为int64类型的再比较
			if n == int64(1) || n == "OK" {
				fmt.Println("add successful")
				// 设置过期时间为24小时
			} else {
				// var maps map[string]string = make(map[string]string)
				fmt.Println("the key has already existed", n)
				fmt.Println("type of n ", reflect.TypeOf(n), reflect.ValueOf(n))

				/*tmp := n.([]interface{})
				for i := 0; i < len(tmp); i = i + 1 {
					// fmt.Printf("tmps is ", tmp)
					fmt.Println("tmps is ", string(tmp[i].([]byte)))
					// maps[string(tmp[i].([]byte))] = string(tmp[i+1].([]byte))
				}*/
				// fmt.Printf("maps res is %#v\n", maps)
			}
		}

	}

}
