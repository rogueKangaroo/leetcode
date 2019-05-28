package simple

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	return reserveByFor(head)
}

func reverseList2(head *ListNode) *ListNode {
	return reserveByRecursion(head)
}

func reserveByFor(l1 *ListNode) *ListNode {
	if l1 == nil || l1.Next == nil {
		return l1
	}
	p := l1
	temp := p.Next
	newTh := temp.Next
	p.Next = nil

	for {
		if newTh == nil {
			temp.Next = p
			break
		}

		temp.Next = p
		p = temp
		temp = newTh
		newTh = newTh.Next
	}
	return temp
}

func reserveByRecursion(l2 *ListNode) *ListNode {
	if l2 == nil || l2.Next == nil {
		return l2
	}
	newTh := reserveByRecursion(l2.Next)
	l2.Next.Next = l2
	l2.Next = nil
	return newTh
}
