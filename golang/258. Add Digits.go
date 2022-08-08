package main

import "fmt"

func digitList(x int) []int {
	list := []int{}
	for x > 0 {
		mod := x % 10
		list = append(list, mod)
		x = x / 10
	}
	return list
}

func addDigits(num int) int {
	if num < 10 {
		return num
	} else {
		digits := digitList(num)
		sum := 0
		for i := 0; i < len(digits); i++ {
			sum += digits[i]
		}

		if sum < 10 {
			return sum
		} else {
			return addDigits(sum)
		}
	}
}

func main() {
	res := addDigits(38)
	fmt.Println(res)
}
