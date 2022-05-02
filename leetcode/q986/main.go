package main

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	answer := make([][]int, 0)
	i := 0
	j := 0
	for {
		for i < len(firstList) && j < len(secondList) && firstList[i][1] < secondList[j][0] {
			i++
		}
		for i < len(firstList) && j < len(secondList) && secondList[j][1] < firstList[i][0] {
			j++
		}
		if i >= len(firstList) || j >= len(secondList) {
			break
		}
		if firstList[i][1] < secondList[j][0] || secondList[j][1] < firstList[i][0] {
			continue
		}
		temp_answer := make([]int, 2)
		if firstList[i][0] > secondList[j][0] {
			temp_answer[0] = firstList[i][0]
		} else {
			temp_answer[0] = secondList[j][0]
		}
		if firstList[i][1] < secondList[j][1] {
			temp_answer[1] = firstList[i][1]
			i++
		} else {
			temp_answer[1] = secondList[j][1]
			j++
		}
		answer = append(answer, temp_answer)
	}

	return answer
}
