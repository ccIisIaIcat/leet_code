package main

import "fmt"

func removeDuplicates(nums []int) int {
	size := len(nums)
	if size == 0{
		return 0
	}
	position := 0
	for i := 0;i<size-1;i++{
		if i == size-1{
			if nums[position] == nums[position+1]{
				nums = nums[:position+1]
			}
		}else{
			if nums[position] == nums[position+1]{
				nums = append(nums[:position+1],nums[position+2:]...)
			}else{
				position += 1
			}
		}
	}
	return position+1
}

func main() {
	test := []int{1,1,2}
	fmt.Println(removeDuplicates(test))
	fmt.Println(test)
}