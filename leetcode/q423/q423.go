package main

import (
	"fmt"
	"strconv"
)

func originalDigits(s string) string {
	tool_map := make(map[byte]int, 0)
	for i := 0; i < len(s); i++ {
		tool_map[s[i]]++
	}
	tool_list := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	answer := make([]int, 0)
	for i := 0; i < 10; {
		signal := 0
		for j := 0; j < len(tool_list[i]); j++ {
			if tool_map[tool_list[i][j]] == 0 {
				signal = 1
				break
			}
		}
		if signal == 1 {
			i++
		} else {
			for j := 0; j < len(tool_list[i]); j++ {
				tool_map[tool_list[i][j]]--
			}
			answer = append(answer, i)
			i++
		}
	}
	answer_s := ""
	for i := 0; i < len(answer); i++ {
		answer_s += strconv.Itoa(answer[i])
	}
	return answer_s
}

func main() {
	aaa := "zeroonetwothreefourfivesixseveneightnine"
	fmt.Println(originalDigits(aaa))
}
