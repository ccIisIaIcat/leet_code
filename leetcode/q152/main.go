package main

import "fmt"

func maxProduct(nums []int) int {
	max := nums[0]
	for i:=1;i<len(nums);i++{
		if maxProduct_1(nums[:i])>=max{
			max = maxProduct_1(nums[:i])
		}
	}
	if max > maxProduct_1(nums){
		return max
	}else{
		return maxProduct_1(nums)
	}
	
}

func maxProduct_1(nums []int) int {
	if len(nums) == 1{
		return nums[0]
	}
	size := len(nums)
	if maxProduct_1(nums[:size-1])*nums[size-1]>0{
		return maxProduct_1(nums[:size-1])*nums[size-1]
	}else if maxProduct_1(nums[:size-1])*nums[size-1]<0{
		return nums[size-1]
	}else{
		return nums[size-1]
	}
}

func maxProduct_2(nums []int) int {
	if len(nums) == 1{
		return nums[0]
	}
	size := len(nums)
	if maxProduct_2(nums[:size-1])*nums[size-1]<0{
		return maxProduct_2(nums[:size-1])*nums[size-1]
	}else if maxProduct_2(nums[:size-1])*nums[size-1]>0{
		return nums[size-1]
	}else{
		return nums[size-1]
	}
}

func main()  {
	test := []int{2,3,-2,4}
	fmt.Println(maxProduct(test))
	
}