package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	answer := new(TreeNode)
	build_tree(nums, answer)
	return answer

}

func build_tree(nums []int, node_now *TreeNode) {
	fmt.Println(nums)
	if len(nums) == 1 {
		node_now.Val = nums[0]
		return
	}
	if len(nums) == 2 {
		node_now.Val = nums[1]
		node_now.Left = new(TreeNode)
		node_now.Left.Val = nums[0]
		return

	}
	mid_index := (len(nums) - 1) / 2
	node_now.Val = nums[mid_index]
	node_now.Left = new(TreeNode)
	node_now.Right = new(TreeNode)
	build_tree(nums[:mid_index], node_now.Left)
	build_tree(nums[mid_index+1:], node_now.Right)

}

func main() {
	a := []int{1, 2}
	sortedArrayToBST(a)
}
