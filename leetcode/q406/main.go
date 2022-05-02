package main

import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	tool_map := make_tool_map(nums2)
	answer := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		if _, ok := tool_map[nums1[i]]; ok {
			answer[i] = tool_map[nums1[i]]
		} else {
			answer[i] = -1
		}
	}
	return answer
}

func make_tool_map(nums2 []int) map[int]int {
	num_stack := make([]int, 0)
	tool_map := make(map[int]int, 0)
	for i := 0; i < len(nums2); i++ {
		if len(num_stack) == 0 {
			num_stack = append(num_stack, nums2[i])
			continue
		} else if num_stack[len(num_stack)-1] > nums2[i] {
			num_stack = append(num_stack, nums2[i])
			continue
		} else {
			for {
				if len(num_stack) == 0 || num_stack[len(num_stack)-1] >= nums2[i] {
					break
				}
				if _, ok := tool_map[num_stack[len(num_stack)-1]]; !ok {
					tool_map[num_stack[len(num_stack)-1]] = nums2[i]
					if len(num_stack) != 0 {
						num_stack = num_stack[:len(num_stack)-1]
					}
				}
			}
			num_stack = append(num_stack, nums2[i])
		}
	}
	for i := 0; i < len(num_stack); i++ {
		tool_map[num_stack[i]] = -1
	}
	return tool_map
}

func main() {
	lll := []int{4, 1, 2}
	aaa := []int{1, 3, 4, 2}
	fmt.Println(nextGreaterElement(lll, aaa))
}
