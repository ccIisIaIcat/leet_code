package main

func lemonadeChange(bills []int) bool {
	tool_map := make(map[int]int, 0)
	tool_map[5] = 0
	tool_map[10] = 0
	for i := 0; i < len(bills); i++ {
		if bills[i] == 5 {
			tool_map[5]++
		} else if bills[i] == 10 {
			tool_map[5]--
			tool_map[10]++
		} else if bills[i] == 20 {
			if tool_map[10] > 0 {
				tool_map[10]--
				tool_map[5]--
			} else {
				tool_map[5] = tool_map[5] - 3
			}
		}
		if tool_map[5] < 0 || tool_map[10] < 0 {
			return false
		}
	}
	return true
}
