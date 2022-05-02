package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var answer []int

func inorderTraversal(root *TreeNode) []int {
	var answer []int
	if root == nil{
		return answer
	}
	ss(root,&answer)
	return answer
}

func ss(root *TreeNode,location *[]int){
	if (*root).Left !=nil{
		ss((*root).Left,location)
	}
	(*location) = append(*location,(*root).Val)
	if (*root).Right !=nil{
		ss((*root).Right,location)
	}
}


func  main()  {

	fmt.Println(answer)
	fmt.Println(answer)
	
	
}