package main

import "fmt"

func sumNums(n int) int {
	ans := 0
	tool_sum(n, &ans)
	fmt.Println(ans)
	return ans
}

func tool_sum(n int, answer *int) bool {
	*answer += n
	return n > 0 && tool_sum(n-1, answer)
}

func main() {
	sumNums(3)
}
