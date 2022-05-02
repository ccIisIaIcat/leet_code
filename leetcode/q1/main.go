package main

import "fmt"


func twoSum(nums []int, target int) []int {
	sum_now := 0
	answer := make([]int,2)
	
	for i:=0;i<len(nums);i++{
		for j:=i+1;j<len(nums);j++{
			sum_now = nums[i]+nums[j]
			if sum_now == target{
				answer[0] = i
				answer[1] = j
				return answer

			}
		}
	}
	return answer


}

func main()  {

	nums_this := []int{2,7,11,15}
	target := 9
	fmt.Println(twoSum(nums_this,target))
	
}