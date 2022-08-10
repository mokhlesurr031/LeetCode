package main

import "fmt"

func sortArrayByParity(nums []int) []int {
	i, j := 0, len(nums)-1

	for i < j {
		if nums[i]%2 != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j -= 1
		} else {
			i += 1
		}
	}
	//fmt.Println(nums)
	return nums

}

func main() {
	arr := []int{3, 1, 2, 4}
	res := sortArrayByParity(arr)
	fmt.Println(res)
}
