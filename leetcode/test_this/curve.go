package main

import "fmt"

func main() {
	aaa := []int{1, 2, 3}
	b := aaa
	for i := 0; i < 4; i++ {
		b[0]++
		fmt.Printf("%p %p %p %p\n", &b, &aaa, &b[0], &aaa[0])
		fmt.Println(b)
		fmt.Println(aaa)
	}

}
