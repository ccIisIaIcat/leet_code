package main

import "fmt"

func findTheDifference(s string, t string) byte {
	original_map := make(map[byte]int, 0)
	for i := 0; i < len(s); i++ {
		original_map[s[i]] += 1
	}
	new_map := make(map[byte]int, 0)
	for i := 0; i < len(t); i++ {
		if new_map[t[i]] == original_map[t[i]] {
			return t[i]
		} else {
			new_map[t[i]] += 1
		}
	}
	return 0

}

func main() {
	a := "string"
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
}
