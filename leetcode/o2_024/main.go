package main

import "fmt"


type ListNode struct {
     Val int
     Next *ListNode
 }

func reverseList(head *ListNode) *ListNode {
	if head == nil{
		return head
	}
	var num []int
	position := head
	for i:=0;;i++{
		if position == nil{
			break
		}else{
			num = append(num,(*position).Val)
			position = (*position).Next
		}
	}
	size := len(num)
	answer := new(ListNode)
	position = answer
	for i := size-1;i>=0;i--{
		(*position).Val = num[i]
		if i >0{
			(*position).Next = new(ListNode)
			position = (*position).Next
		}
	}
	return answer
}

func main()  {
	aa := []int{1,2,3}
	fmt.Println(aa)
	var aaa ListNode
	aaa.Val = 2
	var bbb ListNode
	bbb.Val = 3
	aaa.Next = &bbb
	reverseList(&aaa)
	
}