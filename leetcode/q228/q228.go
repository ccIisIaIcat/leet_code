package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	i := 0
	answer := make([]string, 0)
	for i < len(nums) {
		original_num := i
		for i < len(nums)-1 && nums[i+1] == nums[i]+1 {
			i++
		}
		end_num := i
		if original_num == end_num {
			answer = append(answer, strconv.Itoa(nums[original_num]))
		} else {
			answer = append(answer, strconv.Itoa(nums[original_num])+"->"+strconv.Itoa(nums[end_num]))
		}
		i++
	}
	return answer
}

func main() {
	nums := []int{0, 1, 2, 4, 5, 7}
	fmt.Println(summaryRanges(nums))
}
