//Inserting an element at a given index
package main

import "fmt"

func main() {
	arr := []int{1, 2, 4, 5}
	fmt.Println("length: ", len(arr))

	//insert 7 at index 6
	idx := 2
	val := 3

	//fmt.Println(arr[:idx+1])
	//fmt.Println(arr[idx:])

	//([1,2,4], [4,5])
	arr = append(arr[:idx+1], arr[idx:]...)
	arr[idx] = val

	fmt.Println(arr)

}
