package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sortSentence(s string) string {
	s_arr := strings.Split(s, " ")
	sentence := make([]string, len(s_arr))
	for i := 0; i < len(s_arr); i++ {
		word_ln := len(s_arr[i])
		word_idx_str := string(s_arr[i][word_ln-1])
		word_idx_int, _ := strconv.Atoi(word_idx_str)
		word := string(s_arr[i][:word_ln-1])
		sentence[word_idx_int-1] = word
	}

	ans := ""

	for i := 0; i < len(sentence); i++ {
		if i == len(sentence)-1 {
			ans += sentence[i]
		} else {
			ans += sentence[i] + " "
		}

	}
	//fmt.Println(ans)
	return ans
}

func main() {
	s := "is2 sentence4 This1 a3"
	res := sortSentence(s)
	fmt.Println(res)
}
