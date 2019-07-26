package hard

/**
给定一个二叉树，返回它的 后序 遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [3,2,1]
**/


type PostorderTraversal_TreeNode struct {
      Val int
      Left *PostorderTraversal_TreeNode
      Right *PostorderTraversal_TreeNode
}

func postorderTraversal(root *PostorderTraversal_TreeNode) []int {
	result := make([]int,0)
	if root != nil {
		postorderRange(root, &result)
	}
	return result
}

func postorderRange(root *PostorderTraversal_TreeNode, array *[]int)  {
	if root.Left != nil {
		postorderRange(root.Left,array)
	}
	if root.Right != nil {
		postorderRange(root.Right,array)
	}
	*array = append(*array,root.Val)
}