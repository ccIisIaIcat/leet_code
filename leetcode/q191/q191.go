package main

import (
	"fmt"
)

func hammingWeight(num uint32) int {
	answer := 0
	for i := 0; i < 32; i++ {
		if (1 << i & num) > 0 {
			answer++
		}
	}
	return answer
}

func main() {
	//fmt.Println(hammingWeight(000101))
	fmt.Println(1 << 2)

}
