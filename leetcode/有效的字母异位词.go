package leetcode

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m1 := map[byte]int{}
	for i := 0; i < len(s); i++ {
		m1[s[i]]++
	}
	for i := 0; i < len(s); i++ {
		if m1[t[i]] > 0 {
			m1[t[i]]--
		} else {
			return false
		}
	}
	return true
}
