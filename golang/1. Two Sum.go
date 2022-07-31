package main

func twoSum(nums []int, target int) []int {
	lookup := make(map[int]int)
	ln := len(nums)
	for i := 0; i < ln; i++ {
		flag := target - nums[i]
		_, ok := lookup[flag]
		if ok == false {
			lookup[nums[i]] = i
			continue
		} else {
			return []int{lookup[flag], i}
		}
	}
	return []int{}
}
