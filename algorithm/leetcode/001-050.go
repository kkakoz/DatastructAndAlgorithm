package leetcode

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

// twoSum 001 两数之和
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

// addTwoNumbers 002 两数相加
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

// lengthOfLongestSubstring 03无重复字符的最长子串
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

// 004
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

// 005 最长回文字串 中心线方法
func longestPalindrome(s string) string {
	start, lontest := 0, 0
	for mid := range s {
		left, right := mid, mid
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left, right = left-1, right+1
		}
		if right-left-1 > lontest {
			lontest = right - left - 1
			start = left + 1
		}

		left, right = mid, mid+1
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left, right = left-1, right+1
		}
		if right-left-1 > lontest {
			lontest = right - left - 1
			start = left + 1
		}
	}
	return s[start : start+lontest]
}

// 005 最长回文字串 动态规划
func longestPalindrome2(s string) string {
	isPalindrom := make([][]bool, len(s))
	start, longtest := 0, 0
	for i := range isPalindrom {
		isPalindrom[i] = make([]bool, len(s))
	}
	for i := range s {
		isPalindrom[i][i] = true
	}
	for i := 1; i < len(s); i++ {
		isPalindrom[i][i-1] = true
	}
	for i := len(s) - 1; i > 0; i-- {
		for j := i + 2; j < len(s); j++ {
			isPalindrom[i][j] = isPalindrom[i+1][j-1] && s[i] == s[j]
			if isPalindrom[i][j] && j-i+1 > longtest {
				start = i
				longtest = j - i + 1
			}
		}

	}
	return s[start : start+longtest]
}

// convert 006 Z
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

// reverse 007 整数反转
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

// isPalindrome 回文数
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	str := strconv.Itoa(x)
	length := len(str)
	midLength := length / 2
	for i := 0; i < midLength; i++ {
		if str[i] != str[length-i-1] {
			return false
		}
	}
	return true
}

// isMatch 10 正则表达式匹配
func isMatch(s string, p string) bool {
	return false
}

// maxArea 12 最多水的容器
func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	max := 0
	for {
		if i >= j {
			break
		}
		min := 0
		if height[i] > height[j] {
			min = height[j]
		} else {
			min = height[i]
		}
		res := min * (j - i)
		if res > max {
			max = res
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return max
}

// longestCommonPrefix 14. 最长公共前缀
func longestCommonPrefix(strs []string) string {
	pre := ""
	initPre := false
	for _, v := range strs {
		if !initPre {
			initPre = true
			pre = v
			continue
		}
		for i, b := range pre {
			if i >= len(v) || v[i] != uint8(b) {
				pre = pre[:i]
				break
			}
		}
	}
	return pre
}

//threeSum 15. 三数之和
func threeSum(nums []int) [][]int {
	res := [][]int{}
	if nums == nil || len(nums) < 3 {
		return res
	}
	sort.Ints(nums)
	length := 0
	pre := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] != pre {
			length++
			nums[length] = nums[i]
		}
	}
	nums = nums[0:length]
	fmt.Println(nums)
	return nil
}

// 019 删除链表的倒数第n个结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	first, second := head, dummy
	for i := 0; i < n; i++ {
		first = first.Next
	}
	for ; first != nil; first = first.Next {
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}

//028
func strStr(haystack string, needle string) int {
	if haystack == needle || needle == "" {
		return 0
	}
	l := len(needle)
	for i := range haystack {
		if i+l > len(haystack) {
			return -1
		}
		if haystack[i:i+l] == needle {
			return i
		}
	}
	return -1
}

// 35 搜索插入位置
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left+1 < right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid
		} else {
			left = mid
		}
	}
	if target <= nums[left] {
		return left
	}
	if target <= nums[right] {
		return right
	}
	if target > nums[right] {
		return right + 1
	}
	return 0
}
