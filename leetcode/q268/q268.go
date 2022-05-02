package main

import "fmt"

func missingNumber(nums []int) int {
	tool_map := make(map[int]int, 0)
	for i := 0; i < len(nums)+1; i++ {
		tool_map[i]++
	}
	for i := 0; i < len(nums); i++ {
		tool_map[nums[i]]--
	}
	answer := 0
	for k, v := range tool_map {
		if v == 1 {
			answer = k
		}
	}
	return answer

}

func main() {
	fmt.Println("lalala")
}
