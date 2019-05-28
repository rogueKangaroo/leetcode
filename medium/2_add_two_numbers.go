package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	newTh := &ListNode{}
	head := newTh
	flag := 0

	for {
		if l1 == nil && l2 == nil {
			if flag == 1 {
				newTh.Next = &ListNode{
					Val: 1,
				}
				newTh = newTh.Next
			}
			head = head.Next
			break
		}
		newTh.Next = &ListNode{}
		newTh = newTh.Next
		if l1 == nil {
			value := l2.Val + flag
			newTh.Val = value % 10
			if value >= 10 {
				flag = 1
			} else {
				flag = 0
			}
			l2 = l2.Next
		} else if l2 == nil {
			value := l1.Val + flag
			newTh.Val = value % 10
			if value >= 10 {
				flag = 1
			} else {
				flag = 0
			}
			l1 = l1.Next
		} else {
			value := l1.Val + l2.Val + flag
			newTh.Val = value % 10
			if value >= 10 {
				flag = 1
			} else {
				flag = 0
			}
			l1 = l1.Next
			l2 = l2.Next
		}
	}

	return head
}
