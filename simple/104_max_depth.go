package simple

/*
给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明: 叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。

// 解法：遍历
*/
type TreeNode_maxDepth struct {
     Val int
     Left *TreeNode_maxDepth
     Right *TreeNode_maxDepth
}

func maxDepth(root *TreeNode_maxDepth) int {
	if root == nil {
		return 0
	}
	nodeList := []*TreeNode_maxDepth{root}
	depth := 0
	for {
		if len(nodeList) == 0 {
			return depth
		}
		tempList := make([]*TreeNode_maxDepth,0)
		for _,node := range nodeList {
			if node.Left != nil {
				tempList = append(tempList,node.Left)
			}

			if node.Right != nil {
				tempList = append(tempList,node.Right)
			}
		}
		nodeList = tempList
		depth = depth + 1
	}
}