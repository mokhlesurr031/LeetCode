package main

import "fmt"

func replaceElements(arr []int) []int {
	ans := []int{}
	for i := 0; i < len(arr); i++ {
		high := 0
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > high {
				high = arr[j]
			}
		}
		ans = append(ans, high)
		high = 0
	}
	ans[len(ans)-1] = -1
	//fmt.Println(ans)
	return ans
}

func main() {
	arr := []int{400}
	res := replaceElements(arr)
	fmt.Println(res)
}
