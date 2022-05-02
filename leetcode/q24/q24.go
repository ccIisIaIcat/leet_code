package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	new_head := new(ListNode)
	new_head.Next = head.Next
	head.Next = swapPairs(head.Next.Next)
	head.Next, head.Next.Next = head.Next.Next, head.Next
	return new_head.Next
}

func main() {
	fmt.Println("hello world")
}
