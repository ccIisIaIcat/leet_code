package main

func numJewelsInStones(jewels string, stones string) int {
	tool_map := make(map[byte]bool, 0)
	for i := 0; i < len(jewels); i++ {
		tool_map[jewels[i]] = true
	}
	answer := 0
	for i := 0; i < len(stones); i++ {
		if tool_map[stones[i]] {
			answer++
		}
	}
	return answer
}
