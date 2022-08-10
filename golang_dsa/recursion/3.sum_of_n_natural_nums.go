package main

import "fmt"

func sum_of_n(n int) int {
	if n == 0 {
		return 0
	} else {
		return sum_of_n(n-1) + n
	}
}

func main() {
	n := 100
	sum := sum_of_n(n)
	fmt.Println(sum)
}
