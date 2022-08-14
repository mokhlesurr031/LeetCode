package main

import "fmt"

func romanToInt(s string) int {
	mp := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	ans := 0
	if len(s) == 1 {
		return mp[s[0]]
	}
	for i := 1; i < len(s); i++ {
		if mp[s[i-1]] >= mp[s[i]] {
			ans += mp[s[i-1]]
		} else {
			ans += mp[s[i]] - mp[s[i-1]]
			i++
		}
		if i == len(s)-1 {
			ans += mp[s[i]]
		}
	}
	return ans
}

func main() {
	st := "MCMXCIV"
	res := romanToInt(st)
	fmt.Println(res)
}
