package main

import "fmt"

func main() {
	s := "babab"
	str := []rune(s)

	ln := len(str)
	//fmt.Println(ln)
	flag := true

	for i := 0; i < ln; i++ {
		if str[i] == str[ln-1] {
			ln--
		} else {
			flag = false
			fmt.Println("Not palindrome")
			break
		}
	}
	if flag {
		fmt.Println("Palindrome")
	}

}
