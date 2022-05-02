package main

func numSubarrayProductLessThanK(nums []int, k int) int {
	answer := 0
	for left := 0; left < len(nums); left++ {
		temp_product := 1
		right := left
		for right < len(nums) && temp_product < k {
			temp_product = temp_product * nums[right]
		}
		answer += right - left + 1
	}
	return answer
}
