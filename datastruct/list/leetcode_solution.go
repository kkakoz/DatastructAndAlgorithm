package list


func Leet206ReverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func Leet206ReverseListRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := Leet206ReverseListRecursion(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}