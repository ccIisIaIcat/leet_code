package main

import "fmt"


func merge(nums1 []int, m int, nums2 []int, n int)  {
	answer := make([]int,0)
	a,b := 0,0
	for i:=0;;i++{
		if (a == m){
			answer = append(answer,nums2[b:n]...)
			break
		}
		if (b == n){
			answer = append(answer,nums1[a:m]...)
			break
		}else{
			if nums1[a]<=nums2[b]{
				answer = append(answer,nums1[a])
				a += 1
			}else{
				answer = append(answer,nums2[b])
				b += 1
			}
		}

	}
	copy(nums1,answer)
}

func main()  {
	nums1_ := []int{1,2,3,0,0,0}
	m_ := 3
	nums2_ := []int{2,5,6}
	n_ := 3
	merge(nums1_,m_,nums2_,n_)
	fmt.Println(nums1_)
}