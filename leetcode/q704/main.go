package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (high-low)/2 + low
		num := nums[mid]
		if num == target {
			return mid
		} else if num > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func main() {
	fmt.Println("hello!")
	aa := []int{-1, 0, 3, 5, 9, 12}
	fmt.Println(search(aa, 2))
}
