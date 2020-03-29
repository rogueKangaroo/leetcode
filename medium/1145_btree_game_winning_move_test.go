package medium

import (
	"fmt"
	"testing"
	"time"
)

func Test_btreeGameWinningMove(t *testing.T) {
	root := &btreeGameWinningMove_TreeNode{
		Val:1,
	}
	node2 := &btreeGameWinningMove_TreeNode{
		Val:2,
	}
	node3 := &btreeGameWinningMove_TreeNode{
		Val:3,
	}
	node4 := &btreeGameWinningMove_TreeNode{
		Val:4,
	}
	node5 := &btreeGameWinningMove_TreeNode{
		Val:5,
	}
	node6 := &btreeGameWinningMove_TreeNode{
		Val:6,
	}
	node7 := &btreeGameWinningMove_TreeNode{
		Val:7,
	}
	node8 := &btreeGameWinningMove_TreeNode{
		Val:8,
	}
	node9 := &btreeGameWinningMove_TreeNode{
		Val:9,
	}
	node10 := &btreeGameWinningMove_TreeNode{
		Val:10,
	}
	node11 := &btreeGameWinningMove_TreeNode{
		Val:11,
	}
	root.Left = node2
	root.Right = node3
	node2.Left = node4
	node2.Left = node5
	node3.Left = node6
	node3.Right = node7
	node4.Left = node8
	node4.Right = node9
	node5.Left = node10
	node5.Right = node11

	got := btreeGameWinningMove(root, 11, 3)
	fmt.Println(fmt.Sprintf("btreeGameWinningMove result is %v", got))
	time.Sleep(1 * time.Second)
}


func Test_btreeGameWinningMove1(t *testing.T) {
	root := &btreeGameWinningMove_TreeNode{
		Val:1,
	}
	node2 := &btreeGameWinningMove_TreeNode{
		Val:2,
	}
	node3 := &btreeGameWinningMove_TreeNode{
		Val:3,
	}

	root.Left = node2
	root.Right = node3

	got := btreeGameWinningMove(root, 3, 1)
	fmt.Println(fmt.Sprintf("btreeGameWinningMove result is %v", got))
	time.Sleep(1 * time.Second)
}

func Test_btreeGameWinningMove2(t *testing.T) {
	root := &btreeGameWinningMove_TreeNode{
		Val:5,
	}
	node2 := &btreeGameWinningMove_TreeNode{
		Val:9,
	}
	node3 := &btreeGameWinningMove_TreeNode{
		Val:6,
	}
	node4 := &btreeGameWinningMove_TreeNode{
		Val:7,
	}
	node5 := &btreeGameWinningMove_TreeNode{
		Val:8,
	}
	node6 := &btreeGameWinningMove_TreeNode{
		Val:1,
	}
	node7 := &btreeGameWinningMove_TreeNode{
		Val:4,
	}
	node8 := &btreeGameWinningMove_TreeNode{
		Val:3,
	}
	node9 := &btreeGameWinningMove_TreeNode{
		Val:2,
	}
	root.Left = node2
	root.Right = node3
	node3.Left = node4
	node4.Right = node5
	node5.Left = node6
	node5.Right = node7
	node7.Right = node8
	node8.Right = node9

	got := btreeGameWinningMove(root, 9, 8)
	fmt.Println(fmt.Sprintf("btreeGameWinningMove result is %v", got))
	time.Sleep(1 * time.Second)
}