package simple

import (
	"fmt"
	"testing"
	"time"
)

func Test_isSameTree(t *testing.T) {
	pNode2 := &TreeNode_isSameTree{2,nil,nil}
	pNode3 := &TreeNode_isSameTree{3,nil,nil}
	p := &TreeNode_isSameTree{1,pNode2,pNode3}

	qNode2 := &TreeNode_isSameTree{2,nil,nil}
	qNode3 := &TreeNode_isSameTree{3,nil,nil}
	q := &TreeNode_isSameTree{1,qNode2,qNode3}
	got := isSameTree(p, q)
	fmt.Println(fmt.Sprintf("isSameTree result is %v", got))
	time.Sleep(1 * time.Second)
}


func Test_isSameTree1(t *testing.T) {
	pNode2 := &TreeNode_isSameTree{2,nil,nil}
	p := &TreeNode_isSameTree{1,pNode2,nil}

	qNode2 := &TreeNode_isSameTree{2,nil,nil}
	q := &TreeNode_isSameTree{1,nil,qNode2}
	got := isSameTree(p, q)
	fmt.Println(fmt.Sprintf("isSameTree result is %v", got))
	time.Sleep(1 * time.Second)
}

func Test_isSameTree2(t *testing.T) {
	pNode2 := &TreeNode_isSameTree{2,nil,nil}
	pNode3 := &TreeNode_isSameTree{1,nil,nil}
	p := &TreeNode_isSameTree{1,pNode3,pNode2}

	qNode2 := &TreeNode_isSameTree{2,nil,nil}
	qNode3 := &TreeNode_isSameTree{1,nil,nil}
	q := &TreeNode_isSameTree{1,qNode2,qNode3}
	got := isSameTree(p, q)
	fmt.Println(fmt.Sprintf("isSameTree result is %v", got))
	time.Sleep(1 * time.Second)
}