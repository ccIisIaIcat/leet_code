package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	n_size := len(digits)
	a := make([]int, n_size)
	temp := 0
	if n_size == 1 {
		if digits[0] == 9 {
			a_new := []int{1, 0}
			return a_new
		} else {
			a[0] = digits[0] + 1
			return a
		}

	}
	for i := n_size - 1; i >= 0; i-- {
		if i == n_size-1 {
			temp = digits[i] + 1
			if temp >= 10 {
				a[i] = temp - 10
				digits[i-1]++
			} else {
				a[i] = temp
			}
		} else if i != 0 {
			temp = digits[i]
			if temp >= 10 {
				a[i] = temp - 10
				digits[i-1]++
			} else {
				a[i] = temp
			}
		} else {
			temp = digits[i]
			if temp >= 10 {
				a[i] = temp - 10
				new_a := make([]int, 1)
				new_a[0] = 1
				new_a = append(new_a, a...)
				return new_a
			} else {
				a[i] = temp
			}

		}
	}
	return a
}

func main() {
	fmt.Println("hello world")
	lala := []int{9}
	fmt.Println(plusOne(lala))
}
