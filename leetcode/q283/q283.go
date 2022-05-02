package main

import "fmt"

func moveZeroes(nums []int) {
	i := 0
	n := 0
	for {
		n++
		if n == len(nums) {
			break
		}
		if nums[i] == 0 {
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, 0)
		} else {
			i++
		}

	}
	fmt.Println(nums)
}

func main() {
	aa := []int{0, 0, 1}
	moveZeroes(aa)
}
