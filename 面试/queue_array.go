package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Queue struct {
	Arr  []int
	Max  int
	Head int
	Tail int
}

func Construct(max int) Queue {
	if max <= 0 {
		fmt.Println("queue should more then one.")
	}
	q := Queue{}
	q.Max = max
	q.Arr = make([]int, max)
	return q
}

func (q *Queue) IsNull() bool {
	if q.Head == q.Tail {
		return true
	}
	return false
}

func (q *Queue) IsFull() bool {
	if (q.Tail+1)%q.Max == q.Head {
		return true
	}
	return false
}

func (q *Queue) Length() int {
	if q.Tail >= q.Head {
		return q.Tail - q.Head + 1
	} else {
		return q.Tail + q.Max - q.Head + 1
	}
}

func (q *Queue) Put(value int) {
	if q.IsFull() {
		fmt.Println("queue is full.")
		return
	}
	q.Tail = (q.Tail + 1) % q.Max
	q.Arr[q.Tail] = value
}

func (q *Queue) Get() int {
	if q.IsNull() {
		return -1
	}

	res := q.Arr[q.Head]
	q.Head = (q.Head + 1) % q.Max
	return res
}

func main() {

	q := Construct(20)

	for {
		r := rand.Intn(3)
		if r >= 1 {
			v := rand.Intn(100)
			q.Put(v)
			fmt.Println(r, "PUT", "length:", q.Length(), "value:", v)
		} else {
			t := q.Get()
			fmt.Println(r, "GET", "length:", q.Length(), "value:", t)
		}
		time.Sleep(time.Second * 1)
	}
}
