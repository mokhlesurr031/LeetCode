package main

import (
	"fmt"
)

func checkIfExist(arr []int) bool {
	mp := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		val := arr[i]

		d := val * 2
		h := float64(val) / 2

		is_float := float64(int(h)) != h
		if is_float {
			_, ok_double := mp[d]
			if ok_double {
				return true
			} else {
				mp[val] = val
			}
		} else {
			_, ok_double := mp[d]
			_, ok_half := mp[int(h)]

			if ok_double || ok_half {
				return true
			} else {
				mp[val] = val
			}
		}
	}
	return false
}

func main() {
	arr := []int{10, 2, 5, 3}
	res := checkIfExist(arr)
	fmt.Println(res)
}
