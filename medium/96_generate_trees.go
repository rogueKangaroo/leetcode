package medium

/*

给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？

示例:

输入: 3
输出: 5
解释:
给定 n = 3, 一共有 5 种不同结构的二叉搜索树:

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3

解法：
	dp:
	1.f(i) 代表以i为根节点的树的种数，因为是二叉搜索树，则左子树的数字为[1,i-1]，右子树的数字为[i+1,n]
	2.tree(n) 代表n个数字组成的二叉搜索树的种数，则tree(n) = f(1) + f(2) + ... + f(n)
	3.而由1得，f(i) = tree(i-1)*tree(n-i)
	4.有2和3得：tree(n) = f(1) + f(2) + ... + f(n) = tree(0)tree(n-1) + tree(1)tree(n-2) + ... + tree(n-1)tree(0)
*/
func numTrees(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	tree := make([]int, n+1)
	tree[0] = 1
	tree[1] = 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			tree[i] = tree[i] + tree[j]*tree[i-j-1]
		}
	}
	return tree[n]
}
