package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode{
    var root *ListNode
	for head != nil {
		head.Next, root, head = root, head, head.Next
	}
	return root
	
}


func main()  {
	fmt.Println("he")
	
}