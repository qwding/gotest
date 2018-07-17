package main

import (
	"fmt"
	// 导入redigo扩展包
	"github.com/garyburd/redigo/redis"
	// "strings"
	// "encoding/json"
)

func main() {
	host := "billing-redis-vq1wu.q1.tenxcloud.net:43507"
	db := 0
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
		fmt.Println("please give the command ")
		var cmd string
		fmt.Scanln(&cmd)

		fmt.Println("please give the key.if give q will not use key.")
		var key string
		fmt.Scanln(&key)
		if key == "q" {
			res, err = rs.Do(cmd, key)
			if err != nil {
				fmt.Println("error", err)
				return
			}
			if resArr, ok := res.([]interface{}); ok {
				for j, val := range resArr {
					if str, ok := val.([]byte); ok {
						fmt.Println("res to string is :", j, string(str))
					}
					if strArr, ok := val.([]interface{}); ok {
						for i, v := range strArr {
							fmt.Println("res to array ", j, i, string(v.([]byte)))
						}

					}
				}
			}
			continue
		}

		arr := make([]string, 0)
		for {
			fmt.Println("give q to exit")
			var in string
			fmt.Scanln(&in)

			if in == "q" {
				break
			}
			arr = append(arr, in)
		}
		var res interface{}

		switch len(arr) {
		case 0:
			res, err = rs.Do(cmd, key)
		case 1:
			res, err = rs.Do(cmd, key, arr[0])
		case 2:
			res, err = rs.Do(cmd, key, arr[0], arr[1])
		case 3:
			res, err = rs.Do(cmd, key, arr[0], arr[1], arr[2])
		case 4:
			res, err = rs.Do(cmd, key, arr[0], arr[1], arr[2], arr[3])
		case 5:
			res, err = rs.Do(cmd, key, arr[0], arr[1], arr[2], arr[3], arr[4])
		case 6:
			res, err = rs.Do(cmd, key, arr[0], arr[1], arr[2], arr[3], arr[4], arr[5])
		}

		if err != nil {
			fmt.Println("error", err)
			return
		}
		if resArr, ok := res.([]interface{}); ok {
			for j, val := range resArr {
				if str, ok := val.([]byte); ok {
					fmt.Println("res to string is :", j, string(str))
				}
				if strArr, ok := val.([]interface{}); ok {
					for i, v := range strArr {
						fmt.Println("res to array ", j, i, string(v.([]byte)))
					}

				}
			}
		}

	}
}
