package medium

/*
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n 保证是有效的。

进阶：

你能尝试使用一趟扫描实现吗？
解法：双指针
*/

type ListNode_removeNthFromEnd struct {
	Val int
    Next *ListNode_removeNthFromEnd
}

func removeNthFromEnd(head *ListNode_removeNthFromEnd, n int) *ListNode_removeNthFromEnd {
	if head == nil || (head.Next == nil && n == 1 ){
		return nil
	}
	pointer := head
	index := head
	for {
		if pointer.Next == nil {
			break
		}
		if n==0 {
			index = index.Next
			pointer = pointer.Next
		} else {
			pointer = pointer.Next
			n--
		}
	}
	if n > 0 {
		head = head.Next
	} else if index.Next != nil {
		index.Next = index.Next.Next
	}

	return head
}
