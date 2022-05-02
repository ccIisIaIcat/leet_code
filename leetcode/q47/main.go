package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	answer := make([][]int, 0)
	new_answer := make([][]int, 0)
	answer = append(answer, []int{nums[0]})
	temp_answer := make([]int, 0)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {

		} else {
			for j := 0; j < len(answer); j++ {
				for k := 0; k < len(answer[0])+1; k++ {
					temp_answer = append(answer[j][:k], nums[i])
					temp_answer = append(temp_answer, answer[j][k:]...)
				}
				new_answer = append(new_answer, temp_answer)
			}
		}
		answer = new_answer
	}
	return answer
}

func main() {
	aaa := []int{1, 2, 3, 4}
	fmt.Println(permuteUnique(aaa))
}
