package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	Next *Node
	Data int
}

type Link struct {
	Head *Node
}

func Construct() Link {
	return Link{Head: &Node{}}
}

func (l *Link) Put(data int) {
	node := Node{Data: data}
	node.Next = l.Head.Next
	l.Head.Next = &node
}

func (l *Link) Length() int {
	if l.Head.Next == nil {
		return 0
	}
	n := l.Head.Next
	length := 1
	for n.Next != nil {
		length++
		n = n.Next
	}
	return length
}

func (l *Link) Get() int {
	if l.Head.Next == nil {
		fmt.Println("Link is nil.")
		return -1
	}

	n := l.Head.Next
	for n.Next.Next != nil {
		n = n.Next
	}

	data := n.Next.Data
	n.Next = nil
	return data
}

func main() {

	q := Construct()

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
