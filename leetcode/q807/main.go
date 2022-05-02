package main

import "fmt"

func maxIncreaseKeepingSkyline(grid [][]int) int {
	size := len(grid)
	new_matrix := make([][]int, size)
	for i := 0; i < size; i++ {
		max_point := 0
		for j := 0; j < size; j++ {
			if grid[i][j] > max_point {
				max_point = grid[i][j]
			}
		}
		for j := 0; j < size; j++ {
			new_matrix[i] = append(new_matrix[i], max_point)
		}
	}
	for i := 0; i < size; i++ {
		max_point_2 := 0
		for j := 0; j < size; j++ {
			if grid[j][i] > max_point_2 {
				max_point_2 = grid[j][i]
			}
		}
		for j := 0; j < size; j++ {
			if max_point_2 < new_matrix[j][i] {
				new_matrix[j][i] = max_point_2
			}
		}
	}
	answer := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if grid[i][j] < new_matrix[i][j] {
				answer += new_matrix[i][j] - grid[i][j]
			}
		}
	}
	return answer
}

func main() {
	fmt.Println("hello,world!")
}
