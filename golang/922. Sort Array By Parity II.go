package main

import "fmt"

func sortArrayByParityII(nums []int) []int {
	result := make([]int, len(nums))

	for i := 0; i < 2; i++ {
		idx := i
		if idx == 0 {
			for j := 0; j < len(nums); j++ {
				if nums[j]%2 == 0 {
					result[idx] = nums[j]
					idx += 2
				}
			}
		} else {
			for j := 0; j < len(nums); j++ {
				if nums[j]%2 != 0 {
					result[idx] = nums[j]
					idx += 2
				}
			}
		}
	}
	return result
}

func main() {
	arr := []int{4, 2, 5, 7}
	res := sortArrayByParityII(arr)
	fmt.Println("RES", res)
}
