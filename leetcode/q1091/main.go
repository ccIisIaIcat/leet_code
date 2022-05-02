package main

func shortestPathBinaryMatrix_dfs(grid [][]int) int {
	if grid[0][0] != 0 {
		return -1
	}
	n := len(grid)
	tool_stack := [][]int{{0, 0}}
	temp_position := []int{0, 0}
	answer := 1
	for len(tool_stack) != 0 {
		if temp_position[0] == n-1 && temp_position[1] == n-1 {
			return answer
		}
		if temp_position[0]+1 < n && grid[temp_position[0]+1][temp_position[1]] == 0 {
			temp_position[0]++
			tool_stack = append(tool_stack, temp_position)
			answer++
		} else if temp_position[1]+1 < n && grid[temp_position[0]][temp_position[1]+1] == 0 {
			temp_position[1]++
			tool_stack = append(tool_stack, temp_position)
			answer++
		} else if ((temp_position[1]+1) < n && (temp_position[0]+1) < n) && grid[temp_position[0]+1][temp_position[1]+1] == 0 {
			temp_position[0]++
			temp_position[1]++
			tool_stack = append(tool_stack, temp_position)
			answer++
		} else if ((temp_position[0]+1) < n && (temp_position[1]-1) >= 0) && grid[temp_position[0]+1][temp_position[1]-1] == 0 {
			temp_position[0]++
			temp_position[1]--
			tool_stack = append(tool_stack, temp_position)
			answer++
		} else if temp_position[1]-1 >= 0 && grid[temp_position[0]][temp_position[1]-1] == 0 {
			temp_position[1]--
			tool_stack = append(tool_stack, temp_position)
			answer++
		} else if ((temp_position[0]-1) >= 0 && (temp_position[1]-1) > 0) && grid[temp_position[0]-1][temp_position[1]-1] == 0 {
			temp_position[0]--
			temp_position[1]--
			tool_stack = append(tool_stack, temp_position)
			answer++
		} else if ((temp_position[0] - 1) >= 0) && grid[temp_position[0]-1][temp_position[1]] == 0 {
			temp_position[0]--
			tool_stack = append(tool_stack, temp_position)
			answer++
		} else if ((temp_position[0]-1) >= 0 && temp_position[1]+1 < n) && grid[temp_position[0]-1][temp_position[1]+1] == 0 {
			temp_position[0]--
			temp_position[1]++
			tool_stack = append(tool_stack, temp_position)
			answer++
		} else {
			grid[temp_position[0]][temp_position[1]] = 1
			answer--
			tool_stack = tool_stack[:len(tool_stack)-1]
		}

	}
	return -1

}

func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] != 0 {
		return -1
	}
	n := len(grid)
	tool_queue := [][]int{{0, 0}}
	temp_position := []int{0, 0}
	length := 1
	tool_queue = append(tool_queue, temp_position)
	for len(tool_queue) != 0 {
		temp_position = tool_queue[0]
		length = grid[temp_position[0]][temp_position[1]] + 1
		if temp_position[0] == n-1 && temp_position[1] == n-1 {
			return length
		}
		if temp_position[0]+1 < n && grid[temp_position[0]+1][temp_position[1]] == 0 {
			tool_queue = append(tool_queue, []int{temp_position[0] + 1, temp_position[1]})
			grid[temp_position[0]+1][temp_position[1]] = length
		}
		if temp_position[1]+1 < n && grid[temp_position[0]][temp_position[1]+1] == 0 {
			tool_queue = append(tool_queue, []int{temp_position[0], temp_position[1] + 1})
			grid[temp_position[0]][temp_position[1]+1] = length
		}
		if ((temp_position[1]+1) < n && (temp_position[0]+1) < n) && grid[temp_position[0]+1][temp_position[1]+1] == 0 {
			tool_queue = append(tool_queue, []int{temp_position[0] + 1, temp_position[1] + 1})
			grid[temp_position[0]+1][temp_position[1]+1] = length
		}
		if ((temp_position[0]+1) < n && (temp_position[1]-1) >= 0) && grid[temp_position[0]+1][temp_position[1]-1] == 0 {
			tool_queue = append(tool_queue, []int{temp_position[0] + 1, temp_position[1] - 1})
			grid[temp_position[0]+1][temp_position[1]-1] = length
		}
		if temp_position[1]-1 >= 0 && grid[temp_position[0]][temp_position[1]-1] == 0 {
			tool_queue = append(tool_queue, []int{temp_position[0], temp_position[1] - 1})
			grid[temp_position[0]][temp_position[1]-1] = length
		}
		if ((temp_position[0]-1) >= 0 && (temp_position[1]-1) > 0) && grid[temp_position[0]-1][temp_position[1]-1] == 0 {
			tool_queue = append(tool_queue, []int{temp_position[0] - 1, temp_position[1] - 1})
			grid[temp_position[0]-1][temp_position[1]-1] = length
		}
		if ((temp_position[0] - 1) >= 0) && grid[temp_position[0]-1][temp_position[1]] == 0 {
			tool_queue = append(tool_queue, []int{temp_position[0] - 1, temp_position[1]})
			grid[temp_position[0]-1][temp_position[1]] = length
		}
		if ((temp_position[0]-1) >= 0 && temp_position[1]+1 < n) && grid[temp_position[0]-1][temp_position[1]+1] == 0 {
			tool_queue = append(tool_queue, []int{temp_position[0] - 1, temp_position[1] + 1})
			grid[temp_position[0]-1][temp_position[1]+1] = length
		}
		tool_queue = tool_queue[1:]
	}

	return -1

}
