package main

import "fmt"

func getMaximumGold_test(grid [][]int) int {
	continent_set := make([][][]int, 0)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != 0 {
				start_position := 0
				sub_continent_set := make([][]int, 0)
				sub_continent_set = append(sub_continent_set, []int{i, j})
				for start_position < len(sub_continent_set) {
					grid[sub_continent_set[start_position][0]][sub_continent_set[start_position][1]] = 0
					cross := 0
					i_new := sub_continent_set[start_position][0]
					j_new := sub_continent_set[start_position][1]
					if i_new > 0 {
						if grid[i_new-1][j_new] != 0 {
							sub_continent_set = append(sub_continent_set, []int{i_new - 1, j_new})
							cross++
						}
					}
					if j_new > 0 {
						if grid[i_new][j_new-1] != 0 {
							sub_continent_set = append(sub_continent_set, []int{i_new, j_new - 1})
							cross++
						}
					}
					if i_new < len(grid)-1 {
						if grid[i_new+1][j_new] != 0 {
							sub_continent_set = append(sub_continent_set, []int{i_new + 1, j_new})
							cross++
						}
					}
					if j_new < len(grid[0])-1 {
						if grid[i_new][j_new+1] != 0 {
							sub_continent_set = append(sub_continent_set, []int{i_new, j_new + 1})
							cross++
						}

					}
					sub_continent_set[start_position] = append(sub_continent_set[start_position], cross)
					start_position++
				}
				continent_set = append(continent_set, sub_continent_set)

			}
		}
	}
	fmt.Println(continent_set)
	return 0
}

func main() {
	test_sample := [][]int{{1, 0, 7}, {2, 0, 6}, {3, 4, 5}, {0, 3, 0}, {9, 0, 20}}
	fmt.Println(test_sample)
	getMaximumGold_test(test_sample)
}
