package main

import (
	"fmt"
	"strings"
)

type Node struct {
	value int
	next  *Node
}

func (n Node) String() string {
	return fmt.Sprintf("%d", n.value)
}

type linkedList struct {
	head *Node
	len  int
}

func (l *linkedList) add(value int) {
	newNode := new(Node)
	newNode.value = value

	if l.head == nil {
		l.head = newNode
	} else {
		iterator := l.head
		for iterator.next != nil {
			iterator = iterator.next
		}
		iterator.next = newNode
	}
}

func (l *linkedList) remove(value int) {
	var previous *Node

	for current := l.head; current != nil; current = current.next {
		if current.value == value {
			if l.head == current {
				l.head = current.next
			} else {
				previous.next = current.next
				return
			}
		}
		previous = current
	}
}

func (l linkedList) get(value int) *Node {
	for iterator := l.head; iterator != nil; iterator = iterator.next {
		if iterator.value == value {
			return iterator
		}
	}
	return nil
}

func (l linkedList) String() string {
	sb := strings.Builder{}

	for iterator := l.head; iterator != nil; iterator = iterator.next {
		//fmt.Println(iterator.value, iterator.next)
		sb.WriteString(fmt.Sprintf("%d ", iterator.value))
	}
	return sb.String()
}

func main() {
	fmt.Println("Hello")
	l := linkedList{}
	l.add(1)
	l.add(2)
	l.add(3)
	fmt.Println(l)
	l.add(4)
	fmt.Println(l)
	fmt.Println(l.get(2))
	fmt.Println(l.get(10))
	l.remove(1)
	fmt.Println(l)
	l.remove(4)
	fmt.Println(l)
}
