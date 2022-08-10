package main

import "fmt"

func main() {
	arr := []int{} //empty array

	//append element in the last index
	arr = append(arr, 1)
	arr = append(arr, 2)
	arr = append(arr, 3)
	arr = append(arr, 4)

	fmt.Println("array: ", arr)
	fmt.Println("length: ", len(arr))
}
