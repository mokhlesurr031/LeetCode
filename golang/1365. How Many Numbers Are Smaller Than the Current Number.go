package main

import "fmt"

func smallerNumbersThanCurrent(nums []int) []int {
	result := []int{}
	ln := len(nums)
	for i := 0; i < ln; i++ {
		count := 0

		for j := 0; j < ln; j++ {
			if i == j {
				continue
			} else {
				if nums[j] < nums[i] {
					count += 1
				}
			}
		}
		result = append(result, count)
		count = 0
	}

	return result
}

func main() {
	arr := []int{8, 1, 2, 2, 3}

	res := smallerNumbersThanCurrent(arr)

	fmt.Println(res)
}
