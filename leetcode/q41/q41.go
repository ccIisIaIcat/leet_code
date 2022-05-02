package main

import "fmt"

func firstMissingPositive(nums []int) int {
	tool_map := make(map[int]bool, 0)
	for i := 0; i < len(nums); i++ {
		tool_map[nums[i]] = true
	}
	answer := 1
	for {
		if !tool_map[answer] {
			break
		}
		answer++
	}
	return answer
}

func main() {
	fmt.Println("hello world")
}
