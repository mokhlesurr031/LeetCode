package main

import "fmt"

func validMountainArray(arr []int) bool {
	is_max := false
	//is_valid = false

	ln := len(arr) - 1
	if len(arr) < 3 {
		return false
	}
	i := 0
	for i < ln {
		if i == 0 {
			if arr[i] > arr[i+1] {
				return false
			}
		}
		if is_max == false {
			if arr[i] < arr[i+1] {
				i += 1
				continue
			} else {
				is_max = true
			}
		} else {
			if arr[i] > arr[i+1] {
				i += 1
				continue
			} else {
				return false
			}
		}
	}
	if arr[ln] < arr[ln-1] {
		return true
	} else {
		return false
	}
}

func main() {
	arr := []int{1, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	res := validMountainArray(arr)
	fmt.Println(res)
}
