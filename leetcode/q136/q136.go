package main

func singleNumber(nums []int) int {
	tool_map := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		tool_map[nums[i]] += 1
	}
	answer := 0
	for k, v := range tool_map {
		if v == 1 {
			answer = k
			break
		}
	}
	return answer

}
