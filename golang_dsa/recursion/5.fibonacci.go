package main

import "fmt"

func fibonacci(n int) int {
	if n < 1 {
		return n
	} else if n == 1 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func main() {
	n := 100
	ans := fibonacci(n)
	fmt.Println(ans)
}
