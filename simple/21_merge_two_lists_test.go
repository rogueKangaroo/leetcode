package simple

import (
	"fmt"
	"strconv"
	"testing"
)


//输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4
func Test_mergeTwoLists(t *testing.T) {
	node13 := ListNode_mergeTwoLists{4,nil}
	node12 := ListNode_mergeTwoLists{2,&node13}
	node11 := ListNode_mergeTwoLists{1,&node12}

	node23 := ListNode_mergeTwoLists{4,nil}
	node22 := ListNode_mergeTwoLists{3,&node23}
	node21 := ListNode_mergeTwoLists{1,&node22}
	got := mergeTwoLists(&node11, &node21)

	for {
		if got != nil {
			fmt.Print(strconv.Itoa(got.Val)+"->")
			got = got.Next
		} else {
			break
		}
	}
}

//输入：2
//输出：1
func Test_mergeTwoLists1(t *testing.T) {
	node11 := ListNode_mergeTwoLists{2,nil}
	node21 := ListNode_mergeTwoLists{1,nil}
	got := mergeTwoLists(&node11, &node21)

	for {
		if got != nil {
			fmt.Print(strconv.Itoa(got.Val)+"->")
			got = got.Next
		} else {
			break
		}
	}
}