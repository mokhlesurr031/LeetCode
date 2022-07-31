package main

func isEven(num int) bool {
	count := 0
	for num > 0 {
		num = num / 10
		count++
	}
	if count%2 == 0 {
		return true
	} else {
		return false
	}
}
func findNumbers(nums []int) int {
	res := 0
	for _, val := range nums {
		r := isEven(val)
		if r {
			res++
		}
	}
	return res

}
