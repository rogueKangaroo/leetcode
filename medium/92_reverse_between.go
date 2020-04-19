package medium

/*
反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。

说明:
1 ≤ m ≤ n ≤ 链表长度。

示例:

输入: 1->2->3->4->5->NULL, m = 2, n = 4
输出: 1->4->3->2->5->NULL

解法：双指针，加链表反转2->3->4 -->  4->3->2
925
*/


type ListNode_reverseBetween struct {
     Val int
     Next *ListNode_reverseBetween
}

func reverseBetween(head *ListNode_reverseBetween, m int, n int) *ListNode_reverseBetween {
	node := head
	var flagNode *ListNode_reverseBetween
	mNode := &ListNode_reverseBetween{}
	var reverseNode1 *ListNode_reverseBetween
	var reverseNode2 *ListNode_reverseBetween
	var reverseNode3 *ListNode_reverseBetween

	isStart := true
	for {
		if node == nil {
			break
		}
		if m >1 && n > 1 {
			flagNode = node
			node = node.Next
			m --
			n --
		}else if m == 1 && n > 1 {
			if isStart {
				mNode = node
				reverseNode1 = node
				isStart = false
				node = node.Next
				n --
			} else {
				if reverseNode2 == nil {
					reverseNode2 = node
					node = node.Next
					n --
				}else if reverseNode3 == nil {
					reverseNode3 = node
					reverseNode2.Next = reverseNode1
					reverseNode1 = reverseNode2
					reverseNode2 = reverseNode3
					reverseNode3 = nil
					node = node.Next
					n --
				}
			}

		} else {
			if reverseNode1 == nil {
				break
			} else if reverseNode2 == nil {
				if node.Next != nil {
					mNode.Next = node.Next
				} else {
					mNode.Next = nil
				}
				node.Next = reverseNode1
				if flagNode != nil {
					flagNode.Next = node
				} else {
					head = node
				}
			} else if  reverseNode3 == nil {
				if node.Next != nil {
					mNode.Next = node.Next
				} else {
					mNode.Next = nil
				}
				reverseNode2.Next = reverseNode1
				if flagNode != nil {
					flagNode.Next = node
				} else {
					head = node
				}
				node.Next = reverseNode2
			}
			break
		}
	}
	return head
}


func reverseBetween1(head *ListNode_reverseBetween, m int, n int) *ListNode_reverseBetween {
	if m == n {
		return head
	}
	preHead := &ListNode_reverseBetween{0,head}
	node := preHead
	var tail *ListNode_reverseBetween
	for i:=1;i<=n;i++{
		if i < m {
			node = node.Next
		} else if i == m {
			tail = node.Next
		} else {
			item := tail.Next
			tail.Next = tail.Next.Next
			item.Next = node.Next
			node.Next = item
		}
	}
	return preHead.Next
}



