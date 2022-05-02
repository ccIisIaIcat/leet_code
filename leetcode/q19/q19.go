package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := 0
	tool_head := head
	for {
		length++
		tool_head = tool_head.Next
		if tool_head == nil {
			break
		}
	}
	aim := length - n
	tool_head = head
	sign := 0
	if aim == 0 {
		return tool_head.Next
	} else {
		for {
			sign++
			if sign == aim {
				tool_head.Next = tool_head.Next.Next
				break
			}
			tool_head = tool_head.Next
		}

	}
	return head
}

func search(head *ListNode) {
	for {
		if head == nil {
			break
		}
		fmt.Println(head.Val)
		head = head.Next
	}
}

func main() {
	fmt.Println("hello!")
	aa := new(ListNode)
	aa.Val = 1
	aa.Next = &ListNode{Val: 2}
	aa = removeNthFromEnd(aa, 2)
	search(aa)
}
