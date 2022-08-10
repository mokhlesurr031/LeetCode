package main

import "fmt"

func selectionSort(items []int) []int {
	ln := len(items)
	var temp int
	for i := 0; i < ln-1; i++ {
		min_idx := i
		for j := i + 1; j < ln; j++ {
			if items[min_idx] > items[j] {
				min_idx = j
			}
		}
		temp = items[i]
		items[i] = items[min_idx]
		items[min_idx] = temp
	}
	return items
}

func main() {
	arr := []int{11, 14, 3, 8, 17, 48}
	fmt.Println(selectionSort(arr))
}
