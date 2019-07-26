package medium

import (
	"fmt"
	"testing"
	"time"
)

/**
给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回锯齿形层次遍历如下：

[
  [3],
  [20,9],
  [15,7]
]

**/

func Test_zigzagLevelOrder(t *testing.T) {
	root := &ZigzagLevelOrder_TreeNode{
		Val:3,
	}
	left1 := &ZigzagLevelOrder_TreeNode{
		Val:9,
	}
	right1 := &ZigzagLevelOrder_TreeNode{
		Val:20,
	}
	left2 := &ZigzagLevelOrder_TreeNode{
		Val:15,
	}
	right2 := &ZigzagLevelOrder_TreeNode{
		Val:7,
	}
	root.Left = left1
	root.Right = right1
	right1.Left = left2
	right1.Right = right2
	got := zigzagLevelOrder(root)
	fmt.Println(fmt.Sprintf("zigzagLevelOrder result is %v", got))
	time.Sleep(1 * time.Second)
}
