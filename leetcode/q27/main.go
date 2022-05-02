package main

import "fmt"

func removeElement(nums []int, val int) int {
	size := len(nums)
	location := 0
	for i:= 0;i<size;i++{
		if nums[location] == val && location != size-1{
			nums = append(nums[:location],nums[location+1:]...)

			continue
		}else if nums[location] == val && location == size-1{
			nums = nums[:location]
			continue
		}
		location += 1
		fmt.Println("location is ",location)
		
	}
	fmt.Println(nums,"lalaal")
	return location
}

func main()  {
	test := []int{1,2,3}
	a := test[:2]
	a = append(a,2)
	removeElement(test,1)

	
	
}