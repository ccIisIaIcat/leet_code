package main

import "fmt"


type ListNode struct {
    Val int
    Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var answer *ListNode
	if (*l1).Val <= (*l2).Val{
		answer = l1
		l1 = (*l1).Next
	}else{
		answer = l2
		l2 = (*l2).Next
	}
    var temp *ListNode
    temp = answer
	for i:=0;;i++{
        //fmt.Println("///////////////////")
        //pri_node(answer)
		if l1 == nil{
			 (*temp).Next = l2
             temp = (*temp).Next
			 return answer
		}else if l2 == nil{
			(*temp).Next = l1
            temp = (*temp).Next
			return answer
		}
		if (*l1).Val <= (*l2).Val{
			(*temp).Next = l1
            temp = (*temp).Next
			l1 = (*l1).Next
		}else{
			(*temp).Next = l2
            temp = (*temp).Next
			l2 = (*l2).Next
		}

	}

}

func pri_node(ll *ListNode)  {
	for i:=0;;i++{
		if ll == nil{
			break
		}
		fmt.Println((*ll).Val)
		ll = (*ll).Next
	}
	
}

func main()  {
	fmt.Println("hello world")
	
}