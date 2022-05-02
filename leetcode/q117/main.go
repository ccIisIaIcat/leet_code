package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	tool_list := []*Node{root}
	for len(tool_list) != 0 {
		temp_list := make([]*Node, 0)
		for i := 0; i < len(tool_list); i++ {
			if temp_list[i].Left != nil {
				temp_list = append(temp_list, temp_list[i].Left)
			}
			if temp_list[i].Right != nil {
				temp_list = append(temp_list, temp_list[i].Right)
			}
		}
		for i := 0; i < len(temp_list)-1; i++ {
			temp_list[i].Next = temp_list[i+1]
		}
		tool_list = temp_list
	}
	return root
}
