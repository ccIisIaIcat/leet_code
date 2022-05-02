package main

import "fmt"

func trap(height []int) int {
	water_gather := make([]int, len(height))
	for i := 1; i < len(height)-1; i++ {
		water_gather[i] = mins(find_max_left(&height, i), find_max_right(&height, i)) - height[i]
	}
	answer := 0
	for i := 1; i < len(height)-1; i++ {
		answer += water_gather[i]
	}
	return answer
}

func find_max_left(nums *[]int, position int) int {
	answer := (*nums)[position]
	for i := position; i >= 0; i-- {
		if (*nums)[i] > answer {
			answer = (*nums)[i]
		}
	}
	return answer
}

func find_max_right(nums *[]int, position int) int {
	answer := (*nums)[position]
	for i := position; i < len(*nums); i++ {
		if (*nums)[i] > answer {
			answer = (*nums)[i]
		}
	}
	return answer
}

func mins(a int, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func main() {
	test := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2}

	fmt.Println(trap(test))

}
