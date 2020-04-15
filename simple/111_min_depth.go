package simple

/*
给定一个二叉树，找出其最小深度。

最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

说明: 叶子节点是指没有子节点的节点。

示例:

给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回它的最小深度  2.
*/
type TreeNode_minDepth struct {
	 Val int
	 Left *TreeNode_minDepth
	 Right *TreeNode_minDepth
}

func minDepth(root *TreeNode_minDepth) int {
	return GetMinDepth(root,0)
}

func GetMinDepth(root *TreeNode_minDepth, currentDepth int) int {
	if root == nil {
		return currentDepth
	}
	currentDepth ++
	if root.Left == nil && root.Right != nil{
		return GetMinDepth(root.Right,currentDepth)
	}else if root.Right == nil && root.Left != nil{
		return GetMinDepth(root.Left,currentDepth)
	}else if root.Right != nil && root.Left != nil{
		rightDepth := GetMinDepth(root.Right,currentDepth)
		leftDepth := GetMinDepth(root.Left,currentDepth)
		if leftDepth < rightDepth {
			return leftDepth
		} else {
			return rightDepth
		}
	} else {
		return currentDepth
	}
}