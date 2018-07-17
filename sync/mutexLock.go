package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var lock sync.Mutex
var lockflag bool = false
var signal = false

func main() {
	runtime.GOMAXPROCS(2)
	for i := 0; i < 10; i++ {
		switch i % 3 {
		case 0:
			fmt.Println("main before lock", i)
			lock.Lock()
			fmt.Println("main after lock", i)
			go testLock(i)
			// lock.Unlock()
			//如果unlock是在function外面写的话，那么有点像没有加锁一样，因为主进程加锁完瞬间就解锁了，所以解锁还是要放在func里面的。
			//但是如果unlock放在了func里面，那么在执行这个if的时候，就会阻塞在这个if里，直到这个锁解开了，才会继续执行。
			//生成的结果如下 2
		case 1:
			// lock.Lock()
			fmt.Println("main i%3 == 1 and i is ", i)
			// go testLock(i)
		case 2:
			// lock.Lock()
			// fmt.Println("come into i%3 == 2 and i is ", i)
			// go testLock(i)
			fmt.Println("main before lock", i)
			lock.Lock() //因为在这里lock了。但是因为之前已经将这个方法lock了，还没有解锁，所以这里是锁不上的，
			//所以必须等待锁解开了才能再次给他上锁。看结果1就可以看到，主线程一直走到了上一句停下，直到锁解开，才走到了下一行
			//又因为这是在主线程里，所以主线程就会卡在这里不走，直到锁解开
			fmt.Println("main after lock", i)
			go testLock(i)
			// lock.Unlock()

		default:
			go fmt.Println("main default")
		}

	}
	var str string
	fmt.Scanln(&str)
	fmt.Println("main over")
}

func testLock(i int) {
	lockflag = true
	fmt.Println("process start ", i)
	// lock.Lock()
	timeNow := time.Now()
	fmt.Println(timeNow, "process ", i, " and i%3 is ", i%3)
	time.Sleep(5 * time.Second)
	lockflag = false
	lock.Unlock()

}

/*1
D:\go-testcode\sync>go run mutexLock.go
main before lock 0
main after lock 0
main i%3 == 1 and i is  1
main before lock 2
process start  0
2015-07-17 11:01:48.8149086 +0800 CST process  0  and i%3 is  0
main after lock 2
main before lock 3
process start  2
2015-07-17 11:01:53.822898 +0800 CST process  2  and i%3 is  2
main after lock 3
main i%3 == 1 and i is  4
main before lock 5
process start  3
2015-07-17 11:01:58.8301292 +0800 CST process  3  and i%3 is  0
main after lock 5
main before lock 6
process start  5
2015-07-17 11:02:03.8352085 +0800 CST process  5  and i%3 is  2
main after lock 6
main i%3 == 1 and i is  7
main before lock 8
process start  6
2015-07-17 11:02:08.8433061 +0800 CST process  6  and i%3 is  0
main after lock 8
main before lock 9
process start  8
2015-07-17 11:02:13.8478546 +0800 CST process  8  and i%3 is  2
main after lock 9
process start  9
2015-07-17 11:02:18.8511608 +0800 CST process  9  and i%3 is  0
*/

/*2
come into i%3 == 1 and i is  1
come into this pross 0
2015-07-16 12:05:56.3279637 +0800 CST this pross is in  0  and i%3 is  0
come into this pross 2
2015-07-16 12:06:01.3325726 +0800 CST this pross is in  2  and i%3 is  2
come into i%3 == 1 and i is  4
come into this pross 3
2015-07-16 12:06:06.3340775 +0800 CST this pross is in  3  and i%3 is  0
come into this pross 5
2015-07-16 12:06:11.3377557 +0800 CST this pross is in  5  and i%3 is  2
come into i%3 == 1 and i is  7
come into this pross 6
2015-07-16 12:06:16.3388201 +0800 CST this pross is in  6  and i%3 is  0
come into this pross 8
2015-07-16 12:06:21.3419296 +0800 CST this pross is in  8  and i%3 is  2
come into this pross 9
2015-07-16 12:06:26.3451691 +0800 CST this pross is in  9  and i%3 is  0
*/
