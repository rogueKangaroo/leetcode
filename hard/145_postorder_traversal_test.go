package hard

import (
	"fmt"
	"testing"
	"time"
)

/**
给定一个二叉树，返回它的 后序 遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [3,2,1]
**/


func Test_postorderTraversal(t *testing.T) {
	root := &PostorderTraversal_TreeNode{
		Val:1,
	}
	right := &PostorderTraversal_TreeNode{
		Val:2,
	}
	left := &PostorderTraversal_TreeNode{
		Val:3,
	}
	root.Right = right
	right.Left = left
	got := postorderTraversal(root)
	fmt.Println(fmt.Sprintf("postorderTraversal result(true) is %v", got))
	time.Sleep(100 * time.Millisecond)
}
