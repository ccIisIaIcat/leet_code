package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	answer := make([][]int, 0)
	answer = append(answer, []int{})
	position := 0
	n := 0
	for i := 0; i < len(nums); i++ {
		temp_answer := make([]int, 0)
		if i > 0 && nums[i] == nums[i-1] {
			n = len(answer)
			fmt.Println(answer, temp_answer, position, n)
			for j := position; j < n; j++ {
				temp_answer = make([]int, len(answer[j]))
				copy(temp_answer, answer[j])
				temp_answer = append(temp_answer, nums[i])
				answer = append(answer, temp_answer)
			}
			position = n
		} else {
			n = len(answer)
			fmt.Println(answer, temp_answer, position, n)
			for j := 0; j < n; j++ {
				temp_answer = make([]int, len(answer[j]))
				copy(temp_answer, answer[j])
				temp_answer = append(temp_answer, nums[i])
				answer = append(answer, temp_answer)
			}
			position = n
		}
	}
	return answer
}

func main() {

	aaa := []int{1, 2, 3}
	fmt.Println(subsetsWithDup(aaa))
}
