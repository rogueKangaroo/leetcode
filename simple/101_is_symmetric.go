package simple


/*

给定一个二叉树，检查它是否是镜像对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3
但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3
说明:

如果你可以运用递归和迭代两种方法解决这个问题，会很加分。
*/
type TreeNode_isSymmetric struct {
    Val int
    Left *TreeNode_isSymmetric
    Right *TreeNode_isSymmetric
}

func isSymmetric(root *TreeNode_isSymmetric) bool {
	if root == nil {
		return true
	}
	return floatSymmetric([]*TreeNode_isSymmetric{root})
}

func floatSymmetric(nodeList []*TreeNode_isSymmetric) bool {
	if len(nodeList) == 0 {
		return true
	}
	newNodeList := make([]*TreeNode_isSymmetric,0)
	for i:=0;i<len(nodeList);i++ {
		if nodeList[i] == nil && nodeList[len(nodeList)-1-i] == nil {
			continue
		} else if nodeList[i] != nil && nodeList[len(nodeList)-1-i] == nil {
			return false
		} else if nodeList[i] == nil && nodeList[len(nodeList)-1-i] != nil {
			return false
		}
		if nodeList[i].Val != nodeList[len(nodeList)-1-i].Val {
			return false
		} else {
			newNodeList = append(newNodeList,nodeList[i].Left)
			newNodeList = append(newNodeList,nodeList[i].Right)

		}
	}
	return floatSymmetric(newNodeList)
}