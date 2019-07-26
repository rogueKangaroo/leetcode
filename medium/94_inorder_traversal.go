package medium

/**
给定一个二叉树，返回它的中序 遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,3,2]
**/


type InorderTraversal_TreeNode struct {
     Val int
     Left *InorderTraversal_TreeNode
     Right *InorderTraversal_TreeNode
 }

func inorderTraversal(root *InorderTraversal_TreeNode) []int {
	result := make([]int,0)
	if root != nil {
		middleRange(root,&result)
	}
	return result
}

func middleRange(root *InorderTraversal_TreeNode,array *[]int)  {
	if root.Left != nil {
		middleRange(root.Left,array)
	}
	*array = append(*array,root.Val)
	if root.Right != nil {
		middleRange(root.Right,array)
	}
}