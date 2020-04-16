package medium

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// 输入: head = 1->4->3->2->5->2, x = 3
//输出: 1->2->2->4->3->5
func Test_partition(t *testing.T) {
	node5 := &ListNode_partition{2,nil}
	node4 := &ListNode_partition{5,node5}
	node3 := &ListNode_partition{2,node4}
	node2 := &ListNode_partition{3,node3}
	node1 := &ListNode_partition{4,node2}
	head := &ListNode_partition{1,node1}
	got := partition(head, 3)
	for {
		if got == nil {
			break
		} else {
			fmt.Print(fmt.Sprintf(strconv.Itoa(got.Val)) + "->")
			got = got.Next
		}
	}
	time.Sleep(1 * time.Second)
}


func Test_partition1(t *testing.T) {
	head := &ListNode_partition{}
	got := partition(head, 3)
	for {
		if got == nil {
			break
		} else {
			fmt.Print(fmt.Sprintf(strconv.Itoa(got.Val)) + "->")
			got = got.Next
		}
	}
	time.Sleep(1 * time.Second)
}

func Test_partition2(t *testing.T) {
	head := &ListNode_partition{1,nil}
	got := partition(head, 0)
	for {
		if got == nil {
			break
		} else {
			fmt.Print(fmt.Sprintf(strconv.Itoa(got.Val)) + "->")
			got = got.Next
		}
	}
	time.Sleep(1 * time.Second)
}