package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}
	perpare_position := dummy
	for {
		if head == nil || head.Next == nil {
			break
		}
		if head.Val != head.Next.Val {
			perpare_position = perpare_position.Next
			head = head.Next
		} else {
			for head.Next != nil && head.Val == head.Next.Val {
				head = head.Next
			}
			perpare_position.Next = head.Next
			if head != nil {
				head = head.Next
			}
		}
	}
	return dummy.Next
}
