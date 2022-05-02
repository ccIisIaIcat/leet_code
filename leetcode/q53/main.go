package main

import "fmt"

func maxSubArray(nums []int) int {
    size := len(nums)
    total_max := nums[0]
    for j:=0;j<size;j++{
        if max_list(nums[j:])>total_max{
        total_max = max_list(nums[j:])
        }
    }
    return total_max
}


func max_list(single_list []int) int{
    max_num := single_list[0]
    sum_num := 0
    for i:=0;i<len(single_list);i++{
        sum_num += single_list[i]
        if sum_num > max_num{
            max_num = sum_num
        }
    }
    return max_num
    
}




func main()  {
	test := []int{-2,-1}
	fmt.Println(test)
}