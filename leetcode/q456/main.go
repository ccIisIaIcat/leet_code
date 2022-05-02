package main

import "math"

func find132pattern(nums []int) bool {
	tool_stack := make([]int, 0)
	third := math.MinInt64
	tool_stack = append(tool_stack, nums[len(nums)-1])
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < tool_stack[len(tool_stack)-1] && nums[i] < third {
			return true
		}

		for len(tool_stack) > 0 && nums[i] > tool_stack[len(tool_stack)-1] {
			third = tool_stack[len(tool_stack)-1]
			tool_stack = tool_stack[:len(tool_stack)-1]
		}

		tool_stack = append(tool_stack, nums[i])

	}
	return false
}
