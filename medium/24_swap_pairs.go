package medium

/*
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

 

示例:

给定 1->2->3->4, 你应该返回 2->1->4->3.

*/
type ListNode_swapPairs struct {
    Val int
    Next *ListNode_swapPairs
}

func swapPairs(head *ListNode_swapPairs) *ListNode_swapPairs {
	pointer1 := head
	if pointer1 == nil {
		return nil
	}
	pointer2 := pointer1.Next
	if pointer2 == nil {
		return head
	}
	pointer3 := pointer2.Next
	head = pointer2
	pointer4 := head
	isHead := true

	for {

		pointer2.Next = pointer1
		pointer1.Next = pointer3
		if isHead {
			isHead = false
		} else {
			pointer4.Next = pointer2

		}

		if pointer3 == nil {
			break
		} else {
			pointer4 = pointer1
			pointer1 = pointer3
			pointer2 = pointer1.Next
			if pointer2 == nil {
				break
			}
			pointer3 = pointer2.Next
		}
	}
	return head
}
