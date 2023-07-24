package leetcode

// https://leetcode.cn/problems/jewels-and-stones/
func numJewelsInStones(jewels string, stones string) int {
	m := map[byte]int{}
	for i := 0; i < len(jewels); i++ {
		m[jewels[i]]++
	}
	count := 0
	for i := 0; i < len(stones); i++ {
		if _, ok := m[stones[i]]; ok {
			count++
		}
	}
	return count
}
