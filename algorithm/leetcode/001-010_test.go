package leetcode

import (
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
