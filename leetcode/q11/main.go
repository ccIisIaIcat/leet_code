package main

func maxArea(height []int) int {
	left_point := 0
	right_point := len(height) - 1
	max_area := 0
	for {
		if left_point >= right_point {
			break
		}
		max_area = max_((right_point-left_point)*min_(height[left_point], height[right_point]), max_area)
		if height[left_point] < height[right_point] {
			left_point++
		} else {
			right_point--
		}
	}
	return max_area
}

func max_(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func min_(a int, b int) int {
	if a >= b {
		return b
	} else {
		return a
	}
}
