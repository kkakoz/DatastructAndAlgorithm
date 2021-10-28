package leetcode

//twoSum 001 两数之和
func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		i2, ok := m[target-v]
		if ok {
			return []int{i2, i}
		}
		m[v] = i
	}
	return []int{}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//addTwoNumbers 002 两数相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	var l1Val, l2Val int
	var tempSum int
	var initL3 = false
	tempList := l3
	for {
		if l1 == nil && l2 == nil {
			if tempSum == 1 {
				tempList.Next = &ListNode{Val: tempSum}
				tempList = tempList.Next
			}
			break
		}
		l1Val, l2Val = 0, 0
		if l1 != nil {
			l1Val = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l2Val = l2.Val
			l2 = l2.Next
		}
		l3Val := l1Val + l2Val + tempSum
		tempSum = 0
		if l3Val >= 10 {
			l3Val -= 10
			tempSum = 1
		}
		if !initL3 {
			l3.Val = l3Val
			initL3 = true
		} else {
			tempList.Next = &ListNode{Val: l3Val}
			tempList = tempList.Next
		}
	}
	return l3
}

//lengthOfLongestSubstring 无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	return 0
}