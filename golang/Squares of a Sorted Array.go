package main

import (
	"sort"
)

func sortedSquares(nums []int) []int {
	n := []int{}
	for i := 0; i < len(nums); i++ {
		n = append(n, nums[i]*nums[i])
	}
	sort.Ints(n)
	return n
}

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	sortedSquares(nums)
}
