package main

import "fmt"

func decrypt(code []int, k int) []int {
	answer := make([]int, len(code))
	for i := 0; i < len(code); i++ {
		temp := 0
		if k >= 0 {
			for j := 0; j < k; j++ {
				temp += code[(i+j+1)%len(code)]
			}
		} else {
			for j := k; j < 0; j++ {
				temp += code[(i+j+len(code))%len(code)]
			}
		}

		answer[i] = temp
	}
	return answer

}

func main() {
	fmt.Println("hello world")
	code := []int{2, 4, 9, 3}
	fmt.Println(decrypt(code, -2))
}
