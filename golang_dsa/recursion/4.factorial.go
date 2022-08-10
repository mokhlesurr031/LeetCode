package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return factorial(n-1) * n
	}
}

func main() {
	n := 4
	ans := factorial(n)
	fmt.Println(ans)
}
