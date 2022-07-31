package main

func findMaxConsecutiveOnes(nums []int) int {
	maxCount, currCount := 0, 0

	for _, val := range nums {
		if val == 1 {
			currCount += 1
		} else {
			currCount = 0
		}
		if currCount > maxCount {
			maxCount = currCount
		}
	}
	return maxCount

}
