package main

import (
	"fmt"
	"strconv"
)

func countPoints(rings string) int {
	tool_map_list := make([](map[string]bool), 0)
	for i := 0; i < 10; i++ {
		temp_map := make(map[string]bool, 0)
		tool_map_list = append(tool_map_list, temp_map)
	}
	for i := 0; i < len(rings)/2; i++ {
		id, _ := strconv.Atoi(string(rings[2*i+1]))
		tool_map_list[id][string(rings[2*i])] = true
	}
	answer := 0
	for i := 0; i < len(tool_map_list); i++ {
		if tool_map_list[i]["R"] && tool_map_list[i]["G"] && tool_map_list[i]["B"] {
			answer++
		}
	}

	return answer

}

func main() {
	r := "B0B6G0R6R0R6G9"
	fmt.Println(countPoints(r))

}
