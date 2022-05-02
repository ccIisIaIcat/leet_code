package main

import (
	"fmt"
)

func findCircleNum(isConnected [][]int) int {
	judge_list := make([]bool, len(isConnected))
	answer := 0
	for i := 0; i < len(isConnected); i++ {
		if !judge_list[i] {
			answer++
			judge_list[i] = true

		}
	}
	return answer
}

func dfs(judge_list *[]bool, connect_map [][]int, position int) {
	tool_stack := make([]int, 0)
	tool_stack = append(tool_stack, position)
	for len(tool_stack) != 0 {
		signal := 0
		for i := 0; i < len(connect_map); i++ {
			if connect_map[tool_stack[len(tool_stack)-1]][i] == 1 && !(*judge_list)[i] {
				(*judge_list)[i] = true
				signal = 1
				tool_stack = append(tool_stack, i)
				break
			}
		}
		if signal == 0 {
			tool_stack = tool_stack[:len(tool_stack)-1]
		}
	}

}

func main() {
	aaa := make([]bool, 10)
	fmt.Println(aaa)
}
