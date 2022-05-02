package main

import "fmt"

func solve(board [][]byte) {
	m := len(board)
	n := len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 79 {
				temp_stack := [][]int{{i, j}}
				temp_sum := [][]int{{i, j}}
				signal := 0
				board[i][j] = 88
				for len(temp_stack) != 0 {
					fmt.Println(temp_stack)
					temp_position := []int{temp_stack[len(temp_stack)-1][0], temp_stack[len(temp_stack)-1][1]}
					if temp_position[0] == 0 || temp_position[1] == 0 || temp_position[0] == m-1 || temp_position[1] == n-1 {
						signal = 1
					}
					if temp_position[0]+1 < n && board[temp_position[0]+1][temp_position[1]] == 79 {
						temp_position[0]++
						temp_stack = append(temp_stack, temp_position)
						temp_sum = append(temp_sum, temp_position)
						board[temp_position[0]][temp_position[1]] = 88
					} else if temp_position[1]+1 < m && board[temp_position[0]][temp_position[1]+1] == 79 {
						temp_position[1]++
						temp_stack = append(temp_stack, temp_position)
						temp_sum = append(temp_sum, temp_position)
						board[temp_position[0]][temp_position[1]] = 88
					} else if temp_position[0] > 0 && board[temp_position[0]-1][temp_position[1]] == 79 {
						temp_position[0]--
						temp_stack = append(temp_stack, temp_position)
						temp_sum = append(temp_sum, temp_position)
						board[temp_position[0]][temp_position[1]] = 88
					} else if temp_position[1] > 0 && board[temp_position[0]][temp_position[1]-1] == 79 {
						temp_position[1]--
						temp_stack = append(temp_stack, temp_position)
						temp_sum = append(temp_sum, temp_position)
						board[temp_position[0]][temp_position[1]] = 88
					} else {
						temp_stack = temp_stack[:len(temp_stack)-1]
					}
				}
				if signal == 1 {
					for position_id := 0; position_id < len(temp_sum); position_id++ {
						board[temp_sum[position_id][0]][temp_sum[position_id][1]] = 1
					}
				}
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 1 {
				board[i][j] = 79
			}
		}
	}

}

func main() {
	test := make([][]byte, 1)
	test[0] = []byte{'O'}
	solve(test)
	fmt.Println(test)
}
