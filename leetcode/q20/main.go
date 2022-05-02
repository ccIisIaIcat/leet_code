package main

import "fmt"

func isValid(s string) bool {
	tool_stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if len(tool_stack) == 0 {
			tool_stack = append(tool_stack, s[i])
		} else if s[i] == 40 || s[i] == 123 || s[i] == 91 {
			tool_stack = append(tool_stack, s[i])
		} else {
			if s[i] == 41 {
				if len(tool_stack) > 0 && tool_stack[len(tool_stack)-1] == 40 {
					tool_stack = tool_stack[:len(tool_stack)-1]
				} else {
					return false
				}
			}
			if s[i] == 125 {
				if len(tool_stack) > 0 && tool_stack[len(tool_stack)-1] == 123 {
					tool_stack = tool_stack[:len(tool_stack)-1]
				} else {
					return false
				}
			}
			if s[i] == 93 {
				if len(tool_stack) > 0 && tool_stack[len(tool_stack)-1] == 91 {
					tool_stack = tool_stack[:len(tool_stack)-1]
				} else {
					return false
				}
			}
		}
		fmt.Println(string(tool_stack))
	}
	if len(tool_stack) == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	ss := "([]){"
	fmt.Println(isValid(ss))
}
