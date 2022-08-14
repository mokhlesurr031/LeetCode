package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	mp := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		mp[nums[i]] = true
	}

	res := []int{}

	for i := 1; i <= len(nums); i++ {
		if !mp[i] {
			res = append(res, i)
		}
	}
	return res
}

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	ans := findDisappearedNumbers(nums)
	fmt.Println(ans)
}
