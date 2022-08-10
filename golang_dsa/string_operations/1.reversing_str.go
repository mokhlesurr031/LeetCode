package main

import "fmt"

//two pointer approach
func main() {
	s := "mahin"
	str := []rune(s) //converts char to ascii val

	//fmt.Println(str)

	ln := len(str)
	for i := 0; i < ln; i++ {
		str[i], str[ln-1] = str[ln-1], str[i]
		ln = ln - 1
	}
	res := string(str)
	fmt.Println(res)

}
