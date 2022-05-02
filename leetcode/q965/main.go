package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	return subfunction(root.Val, root)

}

func subfunction(val int, root *TreeNode) bool {
	if root == nil {
		return true
	} else if val != root.Val {
		return false
	}
	return subfunction(val, root.Left) && subfunction(val, root.Right)

}
