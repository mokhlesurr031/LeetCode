package main

import "fmt"

func mergeSort(items []int) []int {
	ln := len(items)

	if ln == 1 {
		return items
	}
	mid := int(ln / 2)
	left, right := make([]int, mid), make([]int, ln-mid)

	for i := 0; i < ln; i++ {
		if i < mid {
			left[i] = items[i]
		} else {
			right[i-mid] = items[i]
		}
	}
	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) (result []int) {
	result = make([]int, len(left)+len(right))
	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}
	return
}

func main() {
	arr := []int{23, 56, 43, 567, 23, 1, 5}
	fmt.Println(mergeSort(arr))
}
