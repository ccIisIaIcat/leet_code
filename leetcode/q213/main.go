package main

import "fmt"

func rob(nums []int) int {
	if len(nums)==1{
		return nums[0]
	}
	a1 := 0
	a2 := 0
	a1 = max(rob_sub(nums[:len(nums)-1]),rob_sub(nums[1:]))
	a2 = max(rob_sub(nums[1:len(nums)-1]),a1)
	return a2
}

func rob_sub(nums []int) int {
	if len(nums) == 0{
		return 0
	}else if len(nums) == 1{
		return nums[0]
	}else if len(nums) == 2{
		return max(nums[0],nums[1])
	}
	dp := make([]int,len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0],nums[1])
	complete_dp(&dp,nums)
	return dp[len(nums)-1]

}
func complete_dp(dp_ *[]int,nums_ []int) {
	for i := 2;i <len(*dp_);i++{
		(*dp_)[i] = max((*dp_)[i-1],(*dp_)[i-2]+nums_[i])
	}
}
func max(a_ int,b_ int) int {
	if a_>=b_{
		return a_
	}else{
		return b_
	}
	
}

func main()  {
	test := []int{1,2,3,1}
	fmt.Println(rob(test))
	
}