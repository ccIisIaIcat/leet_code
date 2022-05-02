package main

func maximalRectangle(matrix [][]byte) int {
	left_max_matrix := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		new_left_max_list := make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			if j == 0 {
				if matrix[i][j] == 49 {
					new_left_max_list[j] = 1
				} else {
					new_left_max_list[j] = 0
				}
			} else {
				if new_left_max_list[j-1] == 48 {
					if matrix[i][j] == 48 {
						new_left_max_list[j] = 0
					} else {
						new_left_max_list[j] = 1
					}
				} else {
					if matrix[i][j] == 48 {
						new_left_max_list[j] = 0
					} else {
						new_left_max_list[j] = new_left_max_list[j-1] + 1
					}
				}
			}
			left_max_matrix[i] = new_left_max_list
		}
	}

	total_max_area := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			max_area := left_max_matrix[i][j]
			min_length := left_max_matrix[i][j]
			for n := i; n >= 0; n-- {
				if left_max_matrix[n][j] < min_length {
					min_length = left_max_matrix[n][j]
				}
				if min_length*(i-n+1) > max_area {
					max_area = min_length * (i - n + 1)
				}
			}
			if max_area > total_max_area {
				total_max_area = max_area
			}
		}
	}

	return total_max_area
}
