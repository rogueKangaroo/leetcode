package medium

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// 输入: 1->2->3->4->5->NULL, m = 2, n = 4
//输出: 1->4->3->2->5->NULL
func Test_reverseBetween(t *testing.T) {
	node4 := &ListNode_reverseBetween{5,nil}
	node3 := &ListNode_reverseBetween{4,node4}
	node2 := &ListNode_reverseBetween{3,node3}
	node1 := &ListNode_reverseBetween{2,node2}
	head := &ListNode_reverseBetween{1,node1}
	got := reverseBetween1(head, 2, 4)
	for {
		if got == nil {
			break
		} else {
			fmt.Print(strconv.Itoa(got.Val) + "->")
			got = got.Next
		}
	}
	time.Sleep(1 * time.Second)
}

// 输入: 5
//输出: 5-null
func Test_reverseBetween1(t *testing.T) {
	head := &ListNode_reverseBetween{5,nil}
	got := reverseBetween1(head, 1, 1)
	for {
		if got == nil {
			break
		} else {
			fmt.Print(strconv.Itoa(got.Val) + "->")
			got = got.Next
		}
	}
	time.Sleep(1 * time.Second)
}


// 输入: 3->5 1,2
//输出: 5->3
func Test_reverseBetween2(t *testing.T) {
	node1 := &ListNode_reverseBetween{5,nil}
	head := &ListNode_reverseBetween{3,node1}
	got := reverseBetween1(head, 1, 2)
	for {
		if got == nil {
			break
		} else {
			fmt.Print(strconv.Itoa(got.Val) + "->")
			got = got.Next
		}
	}
	time.Sleep(1 * time.Second)
}

// 输入: 1->2->3->NULL, m = 1, n = 3
//输出: 3->2->1->NULL
func Test_reverseBetween3(t *testing.T) {
	node2 := &ListNode_reverseBetween{3,nil}
	node1 := &ListNode_reverseBetween{2,node2}
	head := &ListNode_reverseBetween{1,node1}
	got := reverseBetween1(head, 1, 3)
	for {
		if got == nil {
			break
		} else {
			fmt.Print(strconv.Itoa(got.Val) + "->")
			got = got.Next
		}
	}
	time.Sleep(1 * time.Second)
}