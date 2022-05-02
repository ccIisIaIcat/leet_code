package main

import "fmt"


func twoSum(nums []int, target int) []int {
	hash_map := make(map[int]int,0)
	hash_map[nums[0]] = 0
	answer := make([]int,2)
	for i:=1;i<len(nums);i++{
		if hash_map[(target-nums[i])] == 0 && (target-nums[i])!=nums[0]{
			hash_map[nums[i]] = i
		}else{
			answer[0] = hash_map[(target-nums[i])]
			answer[1] = i
			return answer
		}
	}
	return answer
}

func main()  {

	nums_this := []int{2,7,11,15}
	target := 9
	
	fmt.Println(twoSum(nums_this,target))
	
}