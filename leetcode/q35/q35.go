package main

import "fmt"

func searchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		} else {
			continue
		}
	}
	return len(nums)
}

func main() {
	fmt.Println("hello world")
}
