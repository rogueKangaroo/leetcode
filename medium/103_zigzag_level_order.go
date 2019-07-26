package medium

/**
给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回锯齿形层次遍历如下：

[
  [3],
  [20,9],
  [15,7]
]

**/



type ZigzagLevelOrder_TreeNode struct {
	Val int
    Left *ZigzagLevelOrder_TreeNode
    Right *ZigzagLevelOrder_TreeNode
}
func zigzagLevelOrder(root *ZigzagLevelOrder_TreeNode) [][]int {
	resultArray := make([][]int,0)
	if root != nil {
		rootArray := make([]*ZigzagLevelOrder_TreeNode,0)
		rootArray = append(rootArray,root)
		zigzagLevelRange(rootArray,&resultArray,0)
	}
	return resultArray
}

func zigzagLevelRange(node []*ZigzagLevelOrder_TreeNode,orderArray *[][]int, level int)  {
	if len(node) == 0 {
		return
	}
	childNodeArray := make([]*ZigzagLevelOrder_TreeNode,0)
	array := make([]int,0)
	if level%2 == 0 {
		for i:=0;i<len(node);i++ {
			array = append(array,node[i].Val)
			if node[i].Left != nil {
				childNodeArray = append(childNodeArray,node[i].Left)
			}
			if node[i].Right != nil {
				childNodeArray = append(childNodeArray,node[i].Right)
			}
		}
	} else {
		for i:=len(node)-1;i>=0;i-- {
			array = append(array,node[i].Val)
			if node[len(node)-1 - i].Left != nil {
				childNodeArray = append(childNodeArray,node[len(node) - 1 - i].Left)
			}
			if node[len(node)-1 - i].Right != nil {
				childNodeArray = append(childNodeArray,node[len(node)-1 - i].Right)
			}
		}
	}
	*orderArray = append(*orderArray,array)
	zigzagLevelRange(childNodeArray,orderArray,level + 1)
}