package main

import "fmt"

func func1(n int) {
	if n > 0 {
		fmt.Println(n)
		func1(n - 1)
	}
}

func main() {
	x := 5
	func1(x)

}
