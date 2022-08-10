package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}

	//pop last element
	ln := len(arr)
	arr = arr[:ln-1]

	fmt.Println(arr)
}
