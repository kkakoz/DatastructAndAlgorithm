package leetcode

import (
	"math"
	"strconv"
	"strings"
)

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
	m := map[int32]struct{}{}
	maxLen := 0
	r := 0
	for i, c := range s {
		m[c] = struct{}{}
		if r <= i {
			r = i + 1
		}
		for _, next := range s[r:] {
			if _, ok := m[next]; ok {
				break
			}
			r++
			m[next] = struct{}{}
		}
		if len(m) > maxLen {
			maxLen = len(m)
		}
		delete(m, c)
	}
	return maxLen
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	mid := length / 2
	i, j := 0, 0
	preMid := 0
	lastMid := 0
	for i+j <= mid {
		preMid = lastMid
		if i >= len(nums1) {
			lastMid = nums2[j]
			j++
			continue
		}
		if j >= len(nums2) {
			lastMid = nums1[i]
			i++
			continue
		}
		if nums1[i] <= nums2[j] {
			lastMid = nums1[i]
			i++
		} else {
			lastMid = nums2[j]
			j++
		}
	}
	if length%2 == 1 {
		return float64(lastMid)
	} else {
		return float64(preMid+lastMid) / 2
	}
}

//convert 006 Z
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	arrs := make([][]byte, numRows)
	curRow := 0
	down := true
	for _, b := range s {
		arrs[curRow] = append(arrs[curRow], byte(b))
		if curRow+1 == numRows {
			down = false
		}
		if curRow == 0 {
			down = true
		}
		if down {
			curRow++
		} else {
			curRow--
		}
	}
	strBuild := strings.Builder{}
	for _, bs := range arrs {
		strBuild.Write(bs)
	}
	return strBuild.String()
}

//reverse 007 整数反转
func reverse(x int) int {
	res := 0
	for {
		res *= 10
		res += x % 10
		x = x / 10
		if x == 0 {
			if res > math.MaxInt32 || res < -math.MaxInt32 {
				return 0
			}
			return res
		}
	}
}

//isPalindrome 回文数
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	str := strconv.Itoa(x)
	length := len(str)
	midLength := length / 2
	for i := 0; i < midLength; i++ {
		if str[i] != str[length - i - 1] {
			return false
		}
	}
	return true
}

//isMatch 10 正则表达式匹配
func isMatch(s string, p string) bool {
	return false
}

// maxArea 12 最多水的容器
func maxArea(height []int) int {
	return 0
}