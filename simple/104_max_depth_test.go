package simple

import (
	"fmt"
	"testing"
	"time"
)

// 给定二叉树 [3,9,20,null,null,15,7]，
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
//返回它的最大深度 3 。
func Test_maxDepth(t *testing.T) {
	node4 := &TreeNode_maxDepth{7,nil,nil}
	node3 := &TreeNode_maxDepth{15,nil,nil}
	node2 := &TreeNode_maxDepth{20,node3,node4}
	node1 := &TreeNode_maxDepth{9,nil,nil}
	root := &TreeNode_maxDepth{3,node1,node2}
	got := maxDepth(root)
	fmt.Println(fmt.Sprintf("maxDepth result is %v", got))
	time.Sleep(1 * time.Second)
}
