package main

func countPoints(points [][]int, queries [][]int) []int {
	answer := make([]int, len(queries))
	for i := 0; i < len(points); i++ {
		for j := 0; j < len(queries); j++ {
			if ((queries[j][0]-points[i][0])*(queries[j][0]-points[i][0]) + (queries[j][1]-points[i][1])*(queries[j][1]-points[i][1])) <= queries[j][2]*queries[j][2] {
				answer[j] += 1
			}
		}
	}
	return answer
}
