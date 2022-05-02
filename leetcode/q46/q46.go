package main

import (
	"fmt"
	"time"
)

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	} else if len(nums) == 0 {
		aa := make([]int, 0)
		return [][]int{aa}
	} else {
		aa := []int{nums[0]}
		answer := [][]int{aa}
		for i := 0; i < len(nums)-1; i++ {
			answer = tool_per(answer, nums[i+1])
		}
		return answer

	}

}

func tool_per(nums [][]int, number int) [][]int {
	answer := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i])+1; j++ {
			lala := make([]int, len(nums[i]))
			copy(lala, nums[i])
			temp_arr := insert_in_position(lala, number, j)
			answer = append(answer, temp_arr)
			time.Sleep(time.Second)
			fmt.Println(temp_arr)
			fmt.Println(answer)
		}
	}
	return answer
}

func insert_in_position(nums []int, point int, position int) []int {
	answer := append(nums[:position], append([]int{point}, nums[position:]...)...)
	return answer
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	//fmt.Println(permute(a))
	b := permute(a)
	fmt.Println(tool_per(b, 99))
	// fmt.Println(insert_in_position(a, 99, 4))
	// fmt.Println(insert_in_position(a, 99, 2))
	// fmt.Println(insert_in_position(a, 99, 3))
}
