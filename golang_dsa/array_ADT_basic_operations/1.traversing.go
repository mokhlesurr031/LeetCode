package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}

	//traversing array using range
	for idx, val := range arr {
		fmt.Println(idx, val)
	}

	//traversing array using for loop

	ln := len(arr)
	for i := 0; i < ln; i++ {
		fmt.Println(arr[i])
	}
}
