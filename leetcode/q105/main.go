package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	length := len(preorder)
	root := new(TreeNode)
	if length == 0 {
		return root
	}
	root.Val = preorder[0]
	if inorder[0] == preorder[0] {
		root.Left = nil
	} else {
		new_left_list_length := 0
		for i := 0; i < length; i++ {
			if inorder[i] == preorder[0] {
				new_left_list_length = i
			}
		}
		root.Left = buildTree(preorder[1:new_left_list_length+1], inorder[0:new_left_list_length])
	}
	if inorder[length-1] == preorder[0] {
		root.Right = nil
	} else {
		new_right_list_length := 0
		for i := 0; i < length; i++ {
			if inorder[i] == preorder[0] {
				new_right_list_length = length - i - 1
			}
		}
		root.Right = buildTree(preorder[length-new_right_list_length:length], inorder[length-new_right_list_length:length])
	}

	return root
}

// func main() {
// 	my_list := []int{1, 2, 3, 4, 5, 6, 7}
// 	new_list := make([]int, 5)
// 	copy(new_list, my_list[2:])
// 	fmt.Println(new_list)
// }
