package main

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)
	tool_stack := make([]int, 0)
	for i := 0; i < n; i++ {
		answer[i] = -1
	}
	for i := 0; i < 2*n-1; i++ {

		for len(tool_stack) > 0 && nums[i%n] > nums[tool_stack[len(tool_stack)-1]] {
			answer[tool_stack[len(tool_stack)-1]] = nums[i%n]
			tool_stack = tool_stack[:len(tool_stack)-1]
		}
		tool_stack = append(tool_stack, i%n)

	}

	return answer
}
