package main

import "fmt"

func canJump(nums []int) bool {
	if len(nums)==1{
		return true
	}
	reach := 0
	for i:=0;i<len(nums)-1;i++{
		reach = max(reach,i+nums[i])
		if i>= reach && i != len(nums){
			return false
		}
	}
	return true

}

func max(a int,b int)  int{
	if a>=b{
		return a
	}else{
		return b
	}
	
}

func main()  {
	test := []int{3,2,1,1,4}
	fmt.Println(canJump(test))
	
}