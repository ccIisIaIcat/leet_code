package main

import "fmt"

func sortColors(nums []int) {
	tool_map := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		tool_map[nums[i]]++
	}
	for i := 0; i < len(nums); i++ {
		if i < tool_map[0] {
			nums[i] = 0
		} else if i < tool_map[0]+tool_map[1] {
			nums[i] = 1
		} else {
			nums[i] = 2
		}
	}
}

func main() {
	fmt.Println("hello world")
	sample := []int{1}
	sortColors(sample)
	fmt.Println(sample)
}
