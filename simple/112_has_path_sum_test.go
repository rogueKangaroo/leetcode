package simple

import (
	"fmt"
	"testing"
	"time"
)

// 示例: 
//给定如下二叉树，以及目标和 sum = 22，
//
//              5
//             / \
//            4   8
//           /   / \
//          11  13  4
//         /  \      \
//        7    2      1
//返回 true, 因为存在目标和为 22 的根节点到叶子节点的路径 5->4->11->2。
func Test_hasPathSum(t *testing.T) {
	node8 := &TreeNode_hasPathSum{7,nil,nil}
	node7 := &TreeNode_hasPathSum{2,nil,nil}
	node6 := &TreeNode_hasPathSum{1,nil,nil}
	node5 := &TreeNode_hasPathSum{11,node8,node7}
	node4 := &TreeNode_hasPathSum{13,nil,nil}
	node3 := &TreeNode_hasPathSum{4,nil,node6}
	node2 := &TreeNode_hasPathSum{4,node5,nil}
	node1 := &TreeNode_hasPathSum{8,node4,node3}
	root := &TreeNode_hasPathSum{5,node2,node1}
	sum := 22
	got := hasPathSum(root, sum)
	fmt.Println(fmt.Sprintf("hasPathSum result is %v", got))
	time.Sleep(1 * time.Second)
}
