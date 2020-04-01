package simple

import (
	"fmt"
	"testing"
	"time"
)

//示例 1:
//
//给定二叉树 [3,9,20,null,null,15,7]
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
//返回 true 。
func Test_isBalanced(t *testing.T) {
	node4 := &TreeNode_isBalanced{7,nil,nil}
	node3 := &TreeNode_isBalanced{15,nil,nil}
	node2 := &TreeNode_isBalanced{20,node3,node4}
	node1 := &TreeNode_isBalanced{9,nil,nil}
	root := &TreeNode_isBalanced{3,node1,node2}
	got := isBalanced(root)
	fmt.Println(fmt.Sprintf("isBalanced result is %v", got))
	time.Sleep(1 * time.Second)
}

// 示例 2:
//给定二叉树 [1,2,2,3,3,null,null,4,4]
//
//       1
//      / \
//     2   2
//    / \
//   3   3
//  / \
// 4   4
//返回 false 。
func Test_isBalanced1(t *testing.T) {
	node6 := &TreeNode_isBalanced{4,nil,nil}
	node5 := &TreeNode_isBalanced{4,nil,nil}
	node4 := &TreeNode_isBalanced{3,nil,nil}
	node3 := &TreeNode_isBalanced{3,node5,node6}
	node2 := &TreeNode_isBalanced{2,nil,nil}
	node1 := &TreeNode_isBalanced{2,node3,node4}
	root := &TreeNode_isBalanced{1,node1,node2}
	got := isBalanced(root)
	fmt.Println(fmt.Sprintf("isBalanced result is %v", got))
	time.Sleep(1 * time.Second)
}