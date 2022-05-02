package main

func majorityElement(nums []int) int {
	tool_map := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		tool_map[nums[i]]++
		if tool_map[nums[i]] > len(nums)/2 {
			return nums[i]
		}
	}
	return 0
}
