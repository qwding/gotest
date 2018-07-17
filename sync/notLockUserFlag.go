package main

import (
	"fmt"
	"runtime"
	"time"
)

var lockflag bool = false
var signal bool = false

func main() {
	runtime.GOMAXPROCS(2) //如果不加这句话，会发现执行结果是先把for循环执行完，然后再执行go func，于是就只执行了一次testLock（）。
	//所以这句话在写多线程时候还是很必要的
	for i := 0; i < 2000; i++ {
		fmt.Println(time.Now(), "lockflag is ", lockflag, "signal ", signal)
		if !lockflag && signal {
			fmt.Println("come into the if ifififfifififififi")
			lockflag = true
			go testLock(i)
		} else {
			switch i % 3 {
			case 0:
				if !lockflag {
					fmt.Println("come into the !lockflag process")
					lockflag = true
					go testLock(i)
					fmt.Println("0 out")
				} else {
					signal = true
				}

			case 1:
				if !lockflag {
					lockflag = true
					go testLock(i)
				} else {
					signal = true
				}
				// fmt.Println(time.Now(), "come into i%3 == 1 and i is ", i)
			case 2:
				if !lockflag {
					lockflag = true
					go testLock(i)
				} else {
					signal = true
				}

			default:
				go fmt.Println("Come into default")
			}
		}

	}

	var str string
	fmt.Scanln(&str)
	fmt.Println("main over")
}

func testLock(i int) {
	fmt.Println("come into this pross", i)
	timeNow := time.Now()

	time.Sleep(1 * time.Second)
	lockflag = false
	fmt.Println("***********************", timeNow, "this pross is in ", i, " and i%3 is ", i%3)
}
