package simple

/*
给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其自底向上的层次遍历为：

[
  [15,7],
  [9,20],
  [3]
]

解法：遍历
*/
type TreeNode_levelOrderBottom struct {
 	Val int
 	Left *TreeNode_levelOrderBottom
 	Right *TreeNode_levelOrderBottom
}

func levelOrderBottom(root *TreeNode_levelOrderBottom) [][]int {
	if root == nil {
		return [][]int{}
	}
	nodeFloatList := make([][]int,0)
	nodeValueList := []*TreeNode_levelOrderBottom{root}
	nodeList := []int{root.Val}
	nodeFloatList = append(nodeFloatList,nodeList)
	for {
		tempList := make([]*TreeNode_levelOrderBottom,0)
		nodeTempList := make([]int,0)
		for _,node := range nodeValueList {
			if node.Left != nil {
				tempList = append(tempList,node.Left)
				nodeTempList = append(nodeTempList,node.Left.Val)

			}

			if node.Right != nil {
				tempList = append(tempList,node.Right)
				nodeTempList = append(nodeTempList,node.Right.Val)
			}
		}
		nodeValueList = tempList
		if len(nodeValueList) == 0 {
			break
		}
		nodeFloatList = append(nodeFloatList,nodeTempList)
	}
	resNodeFloatList := make([][]int,0)
	for i:=0;i<len(nodeFloatList);i++{
		resNodeFloatList = append(resNodeFloatList,nodeFloatList[len(nodeFloatList)-1-i])
	}
	return resNodeFloatList
}
