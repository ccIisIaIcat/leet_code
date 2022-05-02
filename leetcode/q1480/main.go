package main

func runningSum(nums []int) []int {
	answer := make([]int, len(nums))
	answer[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		answer[i] = answer[i-1] + nums[i]
	}
	return answer
}
