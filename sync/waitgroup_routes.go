package main 


import (
	// "time"
	"sync"
)

func main(){
	var wg sync.WaitGroup
	var urls = []string{
	        "1",
	        "2",
	        "3",
	}

	chans := make(chan string,len(urls))

	for _, url := range urls {
	        wg.Add(1)
	        go func(url string) {
	                defer wg.Done()
	                // time.Sleep(time.Second * 5)
	                chans <- url
	                println("out url:",url)
	        }(url)
	}

	go func(){
		wg.Wait()
		close(chans)
	}()
	for u := range chans {			// 这里一直读到chan关闭，所以会阻塞掉
		println("read chans:",u)
	}

	// time.Sleep(time.Second*6)
	println("end!")
}