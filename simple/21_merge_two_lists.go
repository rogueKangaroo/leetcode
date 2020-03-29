package simple

/*
将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
*/

type ListNode_mergeTwoLists struct {
     Val int
     Next *ListNode_mergeTwoLists
}

func mergeTwoLists(l1 *ListNode_mergeTwoLists, l2 *ListNode_mergeTwoLists) *ListNode_mergeTwoLists {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	currentNode1 := l1
	currentNode2 := l2
	header := &ListNode_mergeTwoLists{l1.Val,nil}
	if currentNode1.Val > currentNode2.Val {
		header = &ListNode_mergeTwoLists{l2.Val,nil}
		currentNode2 = currentNode2.Next
	} else {
		currentNode1 = currentNode1.Next
	}
	pointerNode := header
	for {
		if currentNode1 == nil && currentNode2 == nil {
			break
		}
		if currentNode1 == nil {
			for {
				if currentNode2 != nil {
					pointerNode.Next = currentNode2
					currentNode2 = currentNode2.Next
					pointerNode = pointerNode.Next
				} else {
					break
				}

			}
		} else if currentNode2 == nil {
			for {
				if currentNode1 != nil {
					pointerNode.Next = currentNode1
					currentNode1 = currentNode1.Next
					pointerNode = pointerNode.Next
				} else {
					break
				}

			}
		} else {
			if currentNode1.Val > currentNode2.Val {
				pointerNode.Next = currentNode2
				currentNode2 = currentNode2.Next
				pointerNode = pointerNode.Next
			} else {
				pointerNode.Next = currentNode1
				currentNode1 = currentNode1.Next
				pointerNode = pointerNode.Next
			}
		}
	}
	return header
}
