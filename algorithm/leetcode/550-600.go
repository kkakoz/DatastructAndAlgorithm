package leetcode

import "strings"

// 557 反转字符串中的单词3
func reverseWords3(s string) string {
	strs := strings.Split(s, " ")
	res := strings.Builder{}
	for i, str := range strs {
		for j := len(str) - 1; j >= 0; j-- {
			res.WriteByte(str[j])
		}
		if i != len(strs)-1 {
			res.WriteByte(' ')
		}
	}
	return res.String()
}

// "s'teL ekat edoCteeL tsetnoc "
// "s'teL ekat edoCteeL tsetnoc"

//distributeCandies 575 分糖果
func distributeCandies(candyType []int) int {
	m := make(map[int]struct{}, len(candyType))
	l := len(candyType) / 2
	for _, v := range candyType {
		m[v] = struct{}{}
	}
	if len(m) > l {
		return l
	}
	return len(m)
}
