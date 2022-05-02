package main

import (
	"fmt"
	"math/bits"
)

func hammingDistance(x, y int) int {
	return bits.OnesCount(uint(x ^ y))
}

func main() {
	fmt.Println("hello world")
	fmt.Println(3 ^ 4)
	fmt.Println(uint(3 ^ 4))
}
