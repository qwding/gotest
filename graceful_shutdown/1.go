package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Control struct {
	AllStop bool
	Sigs    chan int
	SigNum  int
}

func newControl(sigNum int) *Control {
	res := &Control{SigNum: sigNum}
	res.Sigs = make(chan int, sigNum)
	return res
}

var ctrl *Control

func main() {
	num := 10
	ctrl = newControl(num)

	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func() {
		defer fmt.Println("out!!!!!!!!!")

		defer close(ctrl.Sigs)

		sig := <-gracefulStop
		fmt.Printf("caught sig: %+v\n", sig)
		println("Wait for 2 second to finish processing")
		ctrl.AllStop = true

		cal := ctrl.SigNum
		t := time.NewTimer(time.Duration(2) * time.Second)
	OUT:
		for {

			select {
			case <-ctrl.Sigs:
				cal = cal - 1
			case <-t.C:
				println("time out")
				break OUT
			}
			if cal <= 0 {
				break OUT
			}
		}

		println("main exit")
		// time.Sleep(2 * time.Second)
		// close(ctrl.Sigs)
		os.Exit(0)
	}()

	for i := 0; i < num; i++ {
		go func(i int) {
			count := 0
			for {
				if ctrl.AllStop {
					println("death ", i, count)
					ctrl.Sigs <- 1
					return
				}

				println(i, count)

				time.Sleep(time.Second * 3)
				// println("out", i)
				count++
			}
		}(i)
	}

	select {}
}

func thread(i int) {
	count := 0
	for {
		if ctrl.AllStop {
			println("death ", i, count)
			ctrl.Sigs <- 1
			return
		}

		println(i, count)

		time.Sleep(time.Second * 3)
		// println("out", i)
		count++
	}
}
