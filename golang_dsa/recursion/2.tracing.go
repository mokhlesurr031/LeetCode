package main

import "fmt"

func func2(n int) {
	if n > 0 {
		func2(n - 1)
		fmt.Println(n)
	}
}

func main() {
	x := 5
	func2(x)
}
