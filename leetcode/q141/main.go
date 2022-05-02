package main

import "fmt"





type ListNode struct {
    Val int
    Next *ListNode
}

func hasCycle(head *ListNode) bool {
    search_map := make(map[*ListNode]int)
	for i:=0;;i++{
		if head == nil{
			break
		}else{
			search_map[head]++
			if search_map[head] ==2{
				return true
			}
			head = (*head).Next
		}
	}
	return false
    
}

func main()  {
	var aa ListNode
	aa.Val = 2
	fmt.Printf("type is:%T",&aa)
	momo := make(map[(*ListNode)]int)
	momo[&aa] = 0
	fmt.Println(momo)
	
}