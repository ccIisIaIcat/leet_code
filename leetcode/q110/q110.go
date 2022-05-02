package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func find_height(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		a := find_height(root.Left)
		b := find_height(root.Right)
		if a == -1 && b == -1 {
			return -1
		} else if a-b > 1 || b-1 > 1 {
			return -1
		} else {
			return max(a, b) + 1
		}
	}
}

func isBalanced(root *TreeNode) bool {
	answer := find_height(root)
	if answer != -1 {
		return true
	} else {
		return false
	}
}

func main() {
	r := new(TreeNode)
	r.Val = 1
	r.Left = &TreeNode{Val: 2}
	r.Left.Left = &TreeNode{Val: 3}
	fmt.Println(isBalanced(r))
}
