package main

import "fmt"

func linearearch(datalist []int, key int) bool {
	for _, item := range datalist {
		if key == item {
			return true
		}

	}
	return false
}

func main() {
	items := []int{95, 78, 46, 58, 45, 86, 99, 251, 320}
	fmt.Println(linearearch(items, 58))
	fmt.Println(linearearch(items, 343))
}
