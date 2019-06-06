package medium

/**
给定一个整数 n，生成所有由 1 ... n 为节点所组成的二叉搜索树。

示例:

输入: 3
输出:
[
  [1,null,3,2],
  [3,2,null,1],
  [3,1,null,null,2],
  [2,1,3],
  [1,null,2,null,3]
]
解释:
以上的输出对应以下 5 种不同结构的二叉搜索树：

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3

解法：
	递归:
	已i为根节点，那么其左子树为1->i的值，右子树为i+1->n的值
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var flag []bool
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	flag = make([]bool,n+1)
	return generate(1, n)
}

func generate( left, right int) ([]*TreeNode) {
	newNodes := []*TreeNode{nil}
	leftBranch := []*TreeNode{}
	rightBranch := []*TreeNode{}

	for i:=left;i<=right;i++ {
		if !flag[i] {
			flag[i] = true
			leftBranch = generate(left,i-1)
			rightBranch = generate(i+1,right)
			for _,left := range leftBranch {
				for _,right := range rightBranch {
					newNode := &TreeNode{
						Val:i,
						Left:left,
						Right:right,
					}
					newNodes = append(newNodes,newNode)
				}
			}
			flag[i] = false
		}
	}
	if len(newNodes) > 1{
		return newNodes[1:]
	}
	return newNodes
}