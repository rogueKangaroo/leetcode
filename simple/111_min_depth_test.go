package simple

import (
	"fmt"
	"testing"
	"time"
)

// 示例:
//
//给定二叉树 [3,9,20,null,null,15,7],
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
//返回它的最小深度  2.
func Test_minDepth(t *testing.T) {
	node4 := &TreeNode_minDepth{15,nil,nil}
	node3 := &TreeNode_minDepth{7,nil,nil}
	node2 := &TreeNode_minDepth{20,node4,node3}
	node1 := &TreeNode_minDepth{9,nil,nil}
	root := &TreeNode_minDepth{7,node1,node2}
	got := minDepth(root)
	fmt.Println(fmt.Sprintf("minDepth result is %v", got))
	time.Sleep(100 * time.Millisecond)
}


func Test_minDepth1(t *testing.T) {
	node1 := &TreeNode_minDepth{9,nil,nil}
	root := &TreeNode_minDepth{7,node1,nil}
	got := minDepth(root)
	fmt.Println(fmt.Sprintf("minDepth result is %v", got))
	time.Sleep(100 * time.Millisecond)
}

// [1,2,3,4,null,null,5]
//    1
//   / \
//  2   3
// /     \
//4       5

func Test_minDepth2(t *testing.T) {
	node4 := &TreeNode_minDepth{5,nil,nil}
	node3 := &TreeNode_minDepth{4,nil,nil}
	node2 := &TreeNode_minDepth{3,nil,node4}
	node1 := &TreeNode_minDepth{2,node3,nil}
	root := &TreeNode_minDepth{1,node1,node2}
	got := minDepth(root)
	fmt.Println(fmt.Sprintf("minDepth result is %v", got))
	time.Sleep(100 * time.Millisecond)
}