package main

import "fmt"

func containsDuplicate(nums []int) bool {
	size := len(nums)
	hash := make(map[int]bool,size)
	for i:=0;i<size;i++{
		if hash[nums[i]] == false{
			hash[nums[i]] = true
		}else{
			return true
		}
	}
	return false
}

func main() {
	test := []int{1,2,3}
	fmt.Println(containsDuplicate(test))
	

	
}