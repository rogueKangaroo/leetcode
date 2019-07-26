package medium

import (
	"fmt"
	"testing"
	"time"
)

/**
给定一个二叉树，返回它的 前序 遍历。

 示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,2,3]
**/

func Test_preorderTraversal(t *testing.T) {
	root := &PreorderTraversal_TreeNode{
		Val:1,
	}
	right := &PreorderTraversal_TreeNode{
		Val:2,
	}
	left := &PreorderTraversal_TreeNode{
		Val:3,
	}
	root.Right = right
	right.Left = left
	got := preorderTraversal(root)
	fmt.Println(fmt.Sprintf("preorderTraversal result is %v", got))
	time.Sleep(1 * time.Second)
}
