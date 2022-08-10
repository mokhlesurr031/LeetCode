package main

import "fmt"

func binarySearch(datalist []int, key int) bool {
	low := 0
	high := len(datalist) - 1

	for low <= high {
		mid := (low + high) / 2

		if datalist[mid] < key {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if low == len(datalist) || datalist[low] != key {
		return false
	}
	return true
}

func main() {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(binarySearch(items, 63))
}
