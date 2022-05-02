package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	tool_map := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		tool_map[i+1] = 0
	}
	for i := 0; i < len(nums); i++ {
		tool_map[nums[i]]++
	}
	answer := make([]int, 0)
	for k, v := range tool_map {
		if v == 0 {
			answer = append(answer, k)
		}
	}
	return answer

}

func main() {
	fmt.Println("hello !!")
}
