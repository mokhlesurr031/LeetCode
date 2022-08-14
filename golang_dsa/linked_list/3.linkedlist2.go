package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
}

func main() {
	//nd := node{data: 25}
	//nd2 := node{data: 35}
	//nd3 := node{data: 45}
	//nd4 := node{data: 55}
	//
	//nd.next = &nd2
	//nd2.next = &nd3
	//nd3.next = &nd4
	//
	//fmt.Println(nd)
	//fmt.Println(nd.next)
	//ll := linkedList{head: &nd}
	//fmt.Println(ll)

	nd := node{data: 20}

	ll := new(linkedList)
	ll.head = &node{data: 25}
	ll.head.next = &node{data: 35}
	ll.head.next.next = &nd

	fmt.Println(ll.head.data)
	fmt.Println(ll.head.next.data)
	fmt.Println(ll.head.next.next.data)

}
