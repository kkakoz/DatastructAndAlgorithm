package leetcode

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