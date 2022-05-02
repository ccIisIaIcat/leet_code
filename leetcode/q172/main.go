package main

import "fmt"

func trailingZeroes(n int) int {
	answer := 0
	for i := 1; ; i++ {
		number := 1
		for j := 1; j <= i; j++ {
			number = number * 5
		}

		if n/(number) == 0 {
			break
		} else {
			answer += n / (number)
		}
	}
	return answer
}

func main() {
	fmt.Println(trailingZeroes(30))
}
