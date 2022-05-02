package main

import (
	"fmt"
	"sort"
	"strconv"
)

//"Bronze Medal","Silver Medal"
func findRelativeRanks(score []int) []string {
	answer := make([]string, len(score))
	tool_ := make([]int, len(score))
	copy(tool_, score)
	sort.Sort(sort.Reverse(sort.IntSlice(tool_)))
	tool_map := make(map[int]int, 0)
	for i := 0; i < len(score); i++ {
		tool_map[score[i]] = i
	}
	for i := 0; i < len(score); i++ {
		if i == 0 {
			answer[tool_map[tool_[i]]] = "Gold Medal"
		} else if i == 1 {
			answer[tool_map[tool_[i]]] = "Silver Medal"
		} else if i == 2 {
			answer[tool_map[tool_[i]]] = "Bronze Medal"
		} else {
			answer[tool_map[tool_[i]]] = strconv.Itoa(i + 1)
		}
	}
	return answer
}

func main() {
	test_sample := []int{10, 3, 8, 9, 4}
	fmt.Println(findRelativeRanks(test_sample))
}
