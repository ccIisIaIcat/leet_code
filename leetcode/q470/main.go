package main

import "fmt"

func rand7() int {

	return 0
}

func rand10() int {
	a := rand7()
	b := 0
	for b != 1 {
		b = rand7()
	}
	return a + b/2
}

func main() {
	fmt.Println(7 / 2)
}
