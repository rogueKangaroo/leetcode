package simple

import (
	"fmt"
	"testing"
	"time"
)

/*
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其自底向上的层次遍历为：

[
  [15,7],
  [9,20],
  [3]
]
*/
func Test_levelOrderBottom(t *testing.T) {
	node4 := &TreeNode_levelOrderBottom{7,nil,nil}
	node3 := &TreeNode_levelOrderBottom{15,nil,nil}
	node2 := &TreeNode_levelOrderBottom{20,node3,node4}
	node1 := &TreeNode_levelOrderBottom{9,nil,nil}
	root := &TreeNode_levelOrderBottom{3,node1,node2}
	got := levelOrderBottom(root)
	fmt.Println(fmt.Sprintf("levelOrderBottom result is %v", got))
	time.Sleep(1 * time.Second)
}
