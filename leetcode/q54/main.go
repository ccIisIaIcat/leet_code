package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	x_min := 0
	y_min := 0
	x_length := len(matrix[0])
	y_length := len(matrix)
	answer := make([]int, x_length*y_length)
	if x_length == 1 {
		answer[0] = matrix[0][0]
		return answer
	}
	for i := 1; i < len(matrix[0]); i++ {
		answer[i-1] = matrix[0][i]
	}
	point_now := []int{0, len(matrix[0]) - 1}
	direction := 1
	for i := len(matrix[0]) - 1; i < len(matrix[0])*len(matrix)-1; i++ {
		if point_now[0] == x_length-1 && point_now[1] == y_min-1 {
			direction = 1
			y_min++
		} else if point_now[0] == x_length-1 && point_now[1] == y_length-1 {
			direction = 2
			x_length--
		} else if point_now[0] == x_min-1 && point_now[1] == y_length-1 {
			direction = 3
			y_length--
		} else if point_now[0] == x_min-1 && point_now[1] == y_min-1 {
			direction = 0
		}
		switch direction {
		case 0:
			point_now[1]++
		case 1:
			point_now[0]++
		case 2:
			point_now[0]--
		case 3:
			point_now[1]--
		}
		fmt.Println(answer, direction, point_now, x_min, x_length, y_min, y_length)
		answer[i] = matrix[point_now[0]][point_now[1]]
	}
	answer[x_length*y_length-1] = point_now[0]
	return answer
}

func main() {
	test_sample := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	spiralOrder(test_sample)

}
