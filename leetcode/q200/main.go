package main

import "fmt"

func numIslands(grid [][]byte) int {
	l := len(grid)
	w := len(grid[0])
	answer := 0
	for i := 0; i < l; i++ {
		for j := 0; j < w; j++ {
			if grid[i][j] == 49 {
				answer++
				dfs(&grid, []int{i, j})
			}
		}
	}
	return answer

}

func dfs(grid *[][]byte, position []int) {
	height := len(*grid)
	width := len((*grid)[0])
	(*grid)[position[0]][position[1]] = 48
	if position[1]+1 < width && (*grid)[position[0]][position[1]+1] == 49 {
		dfs(grid, []int{position[0], position[1] + 1})
	}
	if position[1]-1 >= 0 && (*grid)[position[0]][position[1]-1] == 49 {
		dfs(grid, []int{position[0], position[1] - 1})
	}
	if position[0]+1 < height && (*grid)[position[0]+1][position[1]] == 49 {
		dfs(grid, []int{position[0] + 1, position[1]})
	}
	if position[0]-1 >= 0 && (*grid)[position[0]-1][position[1]] == 49 {
		dfs(grid, []int{position[0] - 1, position[1]})
	}

}

func dfs_2(grid *[][]int, position []int) {
	tool_stack := make([][]int, 0)
	tool_stack = append(tool_stack, position)
	h := len(*grid)
	w := len((*grid)[0])
	for len(tool_stack) != 0 {
		temp_position := tool_stack[len(tool_stack)-1]
		(*grid)[temp_position[0]][temp_position[1]] = 48
		if temp_position[0]+1 < h && (*grid)[temp_position[0]+1][temp_position[1]] == 49 {
			tool_stack = append(tool_stack, []int{temp_position[0] + 1, temp_position[1]})
		} else if temp_position[0]-1 >= 0 && (*grid)[temp_position[0]-1][temp_position[1]] == 49 {
			tool_stack = append(tool_stack, []int{temp_position[0] - 1, temp_position[1]})
		} else if temp_position[1]+1 < w && (*grid)[temp_position[0]][temp_position[1]+1] == 49 {
			tool_stack = append(tool_stack, []int{temp_position[0], temp_position[1] + 1})
		} else if temp_position[1]-1 >= 0 && (*grid)[temp_position[0]][temp_position[1]-1] == 49 {
			tool_stack = append(tool_stack, []int{temp_position[0], temp_position[1] - 1})
		} else {
			tool_stack = tool_stack[:len(tool_stack)-1]
		}
	}
}

func main() {
	a := '1'
	b := '0'
	fmt.Println(byte(a), byte(b))
}
