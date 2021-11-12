package leetcode

import (
	"fmt"
	"log"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testCase := []struct {
		nums   []int
		target int
		res    []int
	}{
		{
			nums:   []int{2, 7, 11, 15},
			target: 9,
			res:    []int{0, 1},
		},
		{
			nums:   []int{0, 4, 3, 0},
			target: 0,
			res:    []int{0, 3},
		},
		{
			nums:   []int{-3, 4, 3, 90},
			target: 0,
			res:    []int{0, 2},
		},
	}
	for _, data := range testCase {
		sum := twoSum(data.nums, data.target)
		for i := range data.res {
			if sum[i] != data.res[i] {
				t.Fatalf("two sum res is %v, is should be %v", sum, data.res)
			}
		}
	}
}

func TestAddTwoNumbers(t *testing.T) {
	testData := []struct {
		l1 *ListNode
		l2 *ListNode
		l3 *ListNode
	}{
		{
			l1: &ListNode{2, &ListNode{4, &ListNode{3, nil}}},
			l2: &ListNode{5, &ListNode{6, &ListNode{4, nil}}},
			l3: &ListNode{7, &ListNode{0, &ListNode{8, nil}}},
		},
	}
	for _, data := range testData {
		res := addTwoNumbers(data.l1, data.l2)
		l3 := data.l3
		for l3 != nil {
			if res.Val != l3.Val {
				log.Fatalf("res = %v, should be %v", res, data.l3)
			}
			res = res.Next
			l3 = l3.Next
		}
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {
	testData := []struct {
		s   string
		len int
	}{
		{"abcabcbb", 3},
		{"bbbb", 1},
		{"pwwkew", 3},
		{"anviaj", 5},
		{"aabaab!bb", 3},
	}
	for _, data := range testData {
		if lengthOfLongestSubstring(data.s) != data.len {
			t.Fatalf("s = %s, len = %d", data.s, lengthOfLongestSubstring(data.s))
		}
	}
}

func TestFindMedianSortedArrays(t *testing.T) {
	testData := []struct {
		num1 []int
		num2 []int
		mid  float64
	}{
		{[]int{1, 3}, []int{2}, 2},
		{[]int{1, 2}, []int{3, 4}, 2.5},
		{[]int{0, 0}, []int{0, 0}, 0},
		{[]int{}, []int{1}, 1},
		{[]int{1, 3}, []int{2, 7}, 2.5},
		{[]int{2}, []int{}, 2},
		{[]int{}, []int{2, 3}, 2.5},
	}
	for _, data := range testData {
		res := findMedianSortedArrays(data.num1, data.num2)
		if res != data.mid {
			t.Fatalf("%v and %v mid should %f, not %f", data.num1, data.num2, data.mid, res)
		}
	}

}

func TestConvert(t *testing.T) {
	testData := []struct {
		s1     string
		number int
		res    string
	}{
		{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
	}
	for _, data := range testData {
		res := convert(data.s1, data.number)
		if res != data.res {
			t.Fatalf("conver %s and %v should %s, not %s", data.s1, data.number, data.res, res)

		}
	}
}

func TestReverse(t *testing.T) {
	testData := []struct{
		param int
		res int
	}{
		{123, 321},
		{120, 21},
	}
	for _, data := range testData {
		res := reverse(data.param)
		if res != data.res {
			t.Fatalf("%d should return %d, not %d", data.param, data.res, res)

		}
	}
}

func TestIsPalindrome(t *testing.T)  {
	palindrome := isPalindrome(1001)
	fmt.Println(palindrome)
}

func TestIsMatch(t *testing.T) {
	isMatch("aab", "c*a*b")
}