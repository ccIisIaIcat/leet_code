package main

import "fmt"

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}
	tool_map := make(map[byte]int, 0)
	state_now := make(map[byte]int, 0)
	answer := make([]int, 0)
	for i := 0; i < len(p); i++ {
		tool_map[p[i]] += 1
		state_now[s[i]] += 1
	}
	if judge(state_now, tool_map) {
		answer = append(answer, 0)
	}
	for i := 0; i < len(s)-len(p); i++ {
		state_now[s[i+len(p)]]++
		state_now[s[i]]--
		if judge(state_now, tool_map) {
			answer = append(answer, i+1)
		}
	}
	return answer
}

func judge(map_s map[byte]int, map_p map[byte]int) bool {
	for k := range map_p {
		if map_p[k] != map_s[k] {
			return false
		}
	}
	return true
}

func main() {
	ll1 := "baa"
	ll2 := "aa"
	fmt.Println(findAnagrams(ll1, ll2))
}
