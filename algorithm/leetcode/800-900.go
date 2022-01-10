package leetcode

func middleNode(head *ListNode) *ListNode {
	i, j := head, head
	for {
		if j == nil {
			if i == nil {
				return nil
			}
			return i
		}
		if j.Next == nil {
			return i
		}
		i = i.Next
		j = j.Next.Next
	}
}
