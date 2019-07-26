package medium

/**
给定一个二叉树，返回它的 前序 遍历。

 示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,2,3]
**/

type PreorderTraversal_TreeNode struct {
     Val int
     Left *PreorderTraversal_TreeNode
     Right *PreorderTraversal_TreeNode
}

func preorderTraversal(root *PreorderTraversal_TreeNode) []int {
	result := make([]int,0)
	if root != nil {
		preorderRange(root,&result)
	}
	return result
}

func preorderRange(root *PreorderTraversal_TreeNode,array *[]int)  {
	*array = append(*array,root.Val)
	if root.Left != nil {
		preorderRange(root.Left,array)
	}
	if root.Right != nil {
		preorderRange(root.Right,array)
	}
}