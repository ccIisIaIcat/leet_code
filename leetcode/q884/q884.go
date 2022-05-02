package main

import "fmt"

func build(s string) string {
	tool_stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] != '#' {
			tool_stack = append(tool_stack, s[i])
		} else if len(tool_stack) > 0 {
			tool_stack = tool_stack[:len(tool_stack)-1]
		}
	}
	return string(tool_stack)
}

func backspaceCompare(s string, t string) bool {
	return build(s) == build(t)
}

func main() {
	fmt.Println(backspaceCompare("ab#c", "ab#c"))
}
