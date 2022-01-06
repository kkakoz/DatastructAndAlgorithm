package leetcode

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 104 二叉树最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	stack := []*TreeNode{}
	stack = append(stack, root)
	ans := 0
	for len(stack) > 0 {
		sz := len(stack)
		for sz > 0 {
			cur := stack[0]
			stack = stack[1:]
			if cur.Left != nil {
				stack = append(stack, cur.Left)
			}
			if cur.Right != nil {
				stack = append(stack, cur.Right)
			}
			sz--
		}
		ans++
	}
	return ans
}

// 125 验证回文串
func isPalindrome125(s string) bool {
	ss := strings.ToLower(s)
	l, r := 0, len(s)-1
	for l < r {
		for !IsVal(ss[l]) && l < r {
			l++
		}
		for !IsVal(ss[r]) && l < r {
			r--
		}
		var left = ss[l]
		var rigth = ss[r]
		var temp = string(ss[l])
		var rr = string(ss[r])
		fmt.Println(left, rigth, temp, rr)
		if l < r && ss[l] != ss[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func IsVal(b byte) bool {
	if (b < 'a' || b > 'z') && (b < 'A' || b > 'Z') && (b < '0' || b > '9') {
		return false
	}
	return true
}
