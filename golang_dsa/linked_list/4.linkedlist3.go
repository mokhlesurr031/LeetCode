package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
}

func (l *linkedList) prepend(value int) {
	newNode := node{data: value}
	if l.head != nil {
		newNode.next = l.head
		l.head = &newNode
	} else {
		l.head = &newNode
	}
}

func (l *linkedList) append(value int) {
	newNode := node{data: value}
	if l.head == nil {
		l.head = &newNode
	} else {
		iterator := l.head
		for l.head.next != nil {
			iterator = iterator.next
		}
		iterator.next = &newNode

	}
}

func main() {
	ll := new(linkedList)
	ll.prepend(5)
	ll.append(1)
	ll.prepend(6)
	fmt.Println(ll.head.data)
	fmt.Println(ll.head.next.data)
	fmt.Println(ll.head.next.next.data)
}
