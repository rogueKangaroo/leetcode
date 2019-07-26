package medium

import (
	"fmt"
	"testing"
	"time"
)

func Test_inorderTraversal(t *testing.T) {
	root := &InorderTraversal_TreeNode{
		Val:1,
	}
	right := &InorderTraversal_TreeNode{
		Val:2,
	}
	left  := &InorderTraversal_TreeNode{
		Val:3,
	}
	root.Right = right
	right.Left = left
	got := inorderTraversal(root)
	fmt.Println(fmt.Sprintf("simplifyPath result is %v" ,got))
	time.Sleep(1 * time.Second)
}
