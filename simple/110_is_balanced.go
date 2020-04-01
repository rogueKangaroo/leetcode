package simple

/*
给定一个二叉树，判断它是否是高度平衡的二叉树。

本题中，一棵高度平衡二叉树定义为：

一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。

示例 1:

给定二叉树 [3,9,20,null,null,15,7]

    3
   / \
  9  20
    /  \
   15   7
返回 true 。

示例 2:

给定二叉树 [1,2,2,3,3,null,null,4,4]

       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
返回 false 。
*/
type TreeNode_isBalanced struct {
	 Val int
     Left *TreeNode_isBalanced
     Right *TreeNode_isBalanced
}

func isBalanced(root *TreeNode_isBalanced) bool {
	if root == nil {
		return true
	}
	isBalance,_ := isBalancedAndDepth(root)
	return isBalance
}

func isBalancedAndDepth(root *TreeNode_isBalanced) (bool,int) {
	if root == nil {
		return true,0
	}
	leftIsBalanced,leftDepth := isBalancedAndDepth(root.Left)
	rightIsBalanced,rightDepth := isBalancedAndDepth(root.Right)
	if leftIsBalanced && rightIsBalanced {
		if leftDepth > rightDepth {
			return leftIsBalanced && rightIsBalanced && (leftDepth-rightDepth<=1),leftDepth + 1
		} else {
			return leftIsBalanced && rightIsBalanced && (rightDepth-leftDepth<=1),rightDepth + 1
		}
	} else {
		if leftDepth > rightDepth {
			return false,leftDepth + 1
		} else {
			return false,rightDepth + 1
		}
	}
}
