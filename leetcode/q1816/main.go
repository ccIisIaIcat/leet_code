package main

import (
	"fmt"
	"strings"
)

func truncateSentence(s string, k int) string {
	tool_arr := strings.Split(s, " ")
	answer := ""
	for i := 0; i < len(tool_arr); i++ {
		answer += tool_arr[i]
		if i != len(tool_arr)-1 {
			answer += " "
		}
	}
	return answer
}

func main() {
	a := "hello world"
	arr := strings.Split(a, " ")
	fmt.Println(arr[0])
}
