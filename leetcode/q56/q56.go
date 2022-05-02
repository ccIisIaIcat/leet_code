package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	tool_arr := make([]int, len(intervals))

	for i := 0; i < len(intervals); i++ {
		tool_arr[i] = intervals[i][0]
	}
	sort.Sort(sort.IntSlice{})
	return nil
}

func conbine(a []int, b []int) []int {
	if a[1] < b[0] || b[1] < a[0] {
		return nil
	} else {
		answer := make([]int, 2)
		answer[0] = min(a[0], b[0])
		answer[1] = max(a[1], b[1])
		return answer
	}
}

func min(a int, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func max(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Println("hello world")
	test_sample := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(test_sample))
}
