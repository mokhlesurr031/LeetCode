package main

import "fmt"

type node struct {
	val  int
	next *node
}

//Pointers:
//	& - address of the variable it is next to
//	if a := 25 and we set b := &a
//	then, value of b holds the memory address of a.
//  and its type is a pointer to an int- *int
//

// * has TWO meaning:
//	1. the * symbol next to a VARIABLE means: get the value of the variable
//		that this pointer is pointing to aka DEREFERENCING.
//	Example:
//		if we set, n := *b
//		the value of n is 25, its the value stored at the memory addredd
//		that b is pointing to.
//
//	2. The * symbol next to a TYPE means: the variable being declared is
//		a pointer and it points to an address holding the type of followed by the *
//
//	Example:
//		var myName *string
//		myName is a variable which holds the memory address of a string variable.

func zero(x *int) {
	fmt.Printf("%T", x)
	*x = 0

}

func main() {
	//var a int
	//a = 25
	//fmt.Printf("Type: %T, val: %d\n", a, a)
	//var b *int
	//b = &a
	//fmt.Printf("Type: %T, val: %d", b, *b)

	x := 5
	zero(&x)
	fmt.Println(x)

}
