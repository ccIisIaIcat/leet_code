package main

import (
	"fmt"
	"math/big"
)

func uniquePaths(m int, n int) int {
	answer := Factoreize(m+n-2) / Factoreize(m-1) / Factoreize(n-1)

	return answer
}

func Factoreize(n int) int {
	answer := 1
	for i := 1; i <= n; i++ {
		answer *= i
		fmt.Println(answer, i)
	}
	fmt.Println(answer)
	return answer
}

func main() {
	fmt.Println("hello world")
	//fmt.Println(uniquePaths(23, 12))
	fmt.Println(Factoreize(25))
	a := big.NewInt(222)
}
