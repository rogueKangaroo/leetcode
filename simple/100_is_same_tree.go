package simple


/*

给定两个二叉树，编写一个函数来检验它们是否相同。

如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。

示例 1:

输入:       1         1
          / \       / \
         2   3     2   3

        [1,2,3],   [1,2,3]

输出: true
示例 2:

输入:      1          1
          /           \
         2             2

        [1,2],     [1,null,2]

输出: false
示例 3:

输入:       1         1
          / \       / \
         2   1     1   2

        [1,2,1],   [1,1,2]

输出: false
*/
type TreeNode_isSameTree struct {
     Val int
     Left *TreeNode_isSameTree
     Right *TreeNode_isSameTree
}

func isSameTree(p *TreeNode_isSameTree, q *TreeNode_isSameTree) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil ||p.Val != q.Val {
		return false
	} else {
		return isSameTree(p.Left,q.Left) && isSameTree(p.Right,q.Right)
	}
}
