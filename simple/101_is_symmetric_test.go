package simple

import (
	"fmt"
	"testing"
	"time"
)

//例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
//
//    1
//   / \
//  2   2
// / \ / \
//3  4 4  3
func Test_isSymmetric(t *testing.T) {
	node6 := &TreeNode_isSymmetric{3,nil,nil}
	node5 := &TreeNode_isSymmetric{4,nil,nil}
	node4 := &TreeNode_isSymmetric{4,nil,nil}
	node3 := &TreeNode_isSymmetric{3,nil,nil}
	node2 := &TreeNode_isSymmetric{2,node5,node6}
	node1 := &TreeNode_isSymmetric{2,node3,node4}
	root := &TreeNode_isSymmetric{1,node1,node2}
	got := isSymmetric(root)
	fmt.Println(fmt.Sprintf("isSymmetric result is %v", got))
	time.Sleep(1 * time.Second)
}


//这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
//
//    1
//   / \
//  2   2
//   \   \
//   3    3
func Test_isSymmetric1(t *testing.T) {
	node6 := &TreeNode_isSymmetric{3,nil,nil}
	node3 := &TreeNode_isSymmetric{3,nil,nil}
	node2 := &TreeNode_isSymmetric{2,nil,node6}
	node1 := &TreeNode_isSymmetric{2,nil,node3}
	root := &TreeNode_isSymmetric{1,node1,node2}
	got := isSymmetric(root)
	fmt.Println(fmt.Sprintf("isSymmetric result is %v", got))
	time.Sleep(1 * time.Second)
}