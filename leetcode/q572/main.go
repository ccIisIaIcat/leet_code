package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if (root == nil && subRoot != nil) || (root != nil && subRoot == nil) {
		return false
	}
	if root.Val == subRoot.Val {
		if judge_two_tree(root, subRoot) {
			return true
		} else {
			return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
		}
	}
	return false
}

func judge_two_tree(root_1 *TreeNode, root_2 *TreeNode) bool {
	if root_1 == nil && root_2 == nil {
		return true
	}
	if (root_1 == nil && root_2 != nil) || (root_1 != nil && root_2 == nil) {
		return false
	}
	if root_1.Val != root_2.Val {
		return false
	}
	return judge_two_tree(root_1.Left, root_2.Left) && judge_two_tree(root_1.Right, root_2.Right)
}

func main() {
	tree_1 := TreeNode{Val: 1}
	tree_2 := TreeNode{Val: 1, Left: &tree_1}
	fmt.Println(isSubtree(&tree_2, &tree_1))
}
