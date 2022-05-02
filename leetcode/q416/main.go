package main

import (
	"fmt"
	"sort"
)

func canPartition(nums []int) bool {
	sort.Ints(nums)

	return true
}

func main() {
	lll := []int{8, 6, 2, 7, 5}
	sort.Ints(lll)
	fmt.Println(lll)
}
