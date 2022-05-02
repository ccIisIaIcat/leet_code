package main

import (
	"fmt"
	"math"
)

func subsets(nums []int) [][]int {
	size := len(nums)
	answer := make([]([]int), int(math.Pow(2, 3)))
	size_now := 0
	// new_list := make([]int, 0)
	for i := 0; i < size; i++ {
		size_now = len(answer)
		for j := 0; j < size_now; j++ {
			// new_list = append(answer[j], nums[i])
			// answer[int(size_now+j)] = append(answer[int(size_now+j)], new_list...)
			// if i == 4 {
			// 	fmt.Println(new_list)
			// 	fmt.Println(answer)
			// }
		}
		fmt.Println(answer)
	}

	return answer
}

func main() {

	test := make([]int, 5)
	test[0] = 1
	test[1] = 2
	test[2] = 3
	test[3] = 4
	test[4] = 5
	answer := subsets(test)
	fmt.Println(answer)
	fmt.Println(int(math.Pow(2, 3)))

}
