package main

import "fmt"

func bubbleSort(items []int) []int {
	ln := len(items)
	//fmt.Println(ln)

	for i := 0; i < ln; i++ {
		for j := 0; j < ln-i-1; j++ {
			if items[j] > items[j+1] {
				items[j], items[j+1] = items[j+1], items[j]
			}
		}
	}
	return items
}
func main() {
	arr := []int{11, 14, 3, 8, 17, 48}
	fmt.Println(bubbleSort(arr))

}
