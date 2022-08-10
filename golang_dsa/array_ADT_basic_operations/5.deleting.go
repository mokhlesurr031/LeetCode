package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("before deleing: ", arr)

	//delete 2 from array
	for idx, val := range arr {
		if val == 2 {
			arr = append(arr[:idx], arr[idx+1:]...)
		}
	}

	fmt.Println("after deleting 2 from array: ", arr)
}
