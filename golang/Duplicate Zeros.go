package main

func duplicateZeros(arr []int) {
	// ln = len(arr)
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 && i+1 < len(arr) {
			arr = append(arr[:i+1], arr[i:len(arr)-1]...)
			i++

		}
	}
}
