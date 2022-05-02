package main

//版本一：dfs
func allPathsSourceTarget(graph [][]int) [][]int {
	answer := make([][]int, 0)
	n := len(graph)
	temp_answer := []int{0}
	for {
		if len(graph[temp_answer[len(temp_answer)-1]]) == 0 {
			if temp_answer[len(temp_answer)-1] == n {
				answer = append(answer, temp_answer)
			}
			temp_answer = temp_answer[:len(temp_answer)-1]
			if len(temp_answer) == 0 {
				break
			} else {
				graph[temp_answer[len(temp_answer)-1]] = graph[temp_answer[len(temp_answer)-1]][1:]
			}
		} else {
			temp_answer = append(temp_answer, graph[temp_answer[len(temp_answer)-1]][0])
		}
	}
	return answer
}
