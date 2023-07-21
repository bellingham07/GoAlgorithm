package leetcode

func canConstruct(ransomNote string, magazine string) bool {
	m1 := map[byte]int{}
	for i := 0; i < len(magazine); i++ {
		m1[magazine[i]]++
	}
	for i := 0; i < len(ransomNote); i++ {
		if m1[ransomNote[i]] > 0 {
			m1[ransomNote[i]]--
		} else {
			return false
		}
	}
	return true
}
