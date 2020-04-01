package simple

/*
将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

示例:

给定有序数组: [-10,-3,0,5,9],

一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：

      0
     / \
   -3   9
   /   /
 -10  5

解法：二叉搜索树是指左节点比根节点小，右结点比根节点大
*/
type TreeNode_sortedArrayToBST struct {
     Val int
     Left *TreeNode_sortedArrayToBST
     Right *TreeNode_sortedArrayToBST
}

func sortedArrayToBST(nums []int) *TreeNode_sortedArrayToBST {
	if len(nums) == 0 {
		return nil
	} else if len(nums) == 1 {
		return &TreeNode_sortedArrayToBST{nums[0],nil,nil}
	}
	return sortArrayToBST(nums,0,len(nums) - 1)

}

func sortArrayToBST(nums []int,start,end int) *TreeNode_sortedArrayToBST {
	if start > end {
		return nil
	}
	nodeIndex := start + (end-start+1)/2
	node := &TreeNode_sortedArrayToBST{nums[nodeIndex],nil,nil}
	node.Left = sortArrayToBST(nums,start,nodeIndex-1)
	node.Right = sortArrayToBST(nums,nodeIndex+1,end)
	return node
}