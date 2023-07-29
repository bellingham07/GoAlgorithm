package leetcode

func strStr(haystack string, needle string) int {
	check := func(c string) bool {
		for i := 0; i < len(needle); i++ {
			if needle[i] != c[i] {
				return false
			}
		}
		return true
	}
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i] == needle[0] {
			if check(haystack[i : i+len(needle)]) {
				return i
			}
		}
	}
	return -1
}
