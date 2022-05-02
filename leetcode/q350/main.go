package main

import "fmt"


func intersect(nums1 []int, nums2 []int) []int {
	size1,size2 := len(nums1),len(nums2)
	ori_map := make(map[int]int,0)
	answer := make([]int,0)
	if size1<=size2{
		for i:=0;i<size1;i++{
			ori_map[nums1[i]] += 1
		}
		for j:=0;j<size2;j++{
			if ori_map[nums2[j]]!=0{
				answer = append(answer,nums2[j])
				ori_map[nums2[j]] -= 1
			}
		}
	}
	if size1>size2{
		for i:=0;i<size2;i++{
			ori_map[nums2[i]] += 1
		}
		for j:=0;j<size1;j++{
			if ori_map[nums1[j]]!=0{
				answer = append(answer,nums1[j])
				ori_map[nums1[j]] -= 1
			}
		}
	}
	return answer

}


func main()  {
	nums1 := []int{2}
	nums2 := []int{2}
	fmt.Println(intersect(nums2,nums1))
	
}