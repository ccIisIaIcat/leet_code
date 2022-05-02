package main

func minSubArrayLen(target int, nums []int) int {
	left := 0
	right := 0
	temp_sum := 0
	for temp_sum < target && right < len(nums) {
		temp_sum += nums[right]
		right++
	}
	answer := right - left + 1
	if answer == len(nums) {
		return 0
	}
	for right < len(nums) && left <= right {
		if temp_sum >= target {
			if right-left+1 < answer {
				answer = right - left + 1
			}
			temp_sum -= nums[left]
			left++
		} else {
			temp_sum += nums[right]
			right++
		}

	}

	return answer
}
