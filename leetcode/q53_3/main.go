package main

func maxSubArray(nums []int) int {
	n := len(nums)
	max_int := nums[0]
	for i := 1; i < n; i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		if nums[i] > max_int {
			max_int = nums[i]
		}
	}

	return max_int
}
