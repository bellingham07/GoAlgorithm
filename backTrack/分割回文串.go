package backTrack

var (
	path []string // 放已经回文的子串
	res  [][]string
)

func partition(s string) [][]string {
	path, res = make([]string, 0), make([][]string, 0)
	dfs(s, 0)
	return res
}

func dfs(s string, start int) {
	if start == len(s) { // 如果起始位置等于s的大小，说明已经找到了一组分割方案了
		tmp := make([]string, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		return
	}
	for i := start; i < len(s); i++ {
		str := s[start : i+1]
		if isPalindrome(str) { // 是回文子串
			path = append(path, str)
			dfs(s, i+1)               // 寻找i+1为起始位置的子串
			path = path[:len(path)-1] // 回溯过程，弹出本次已经添加的子串
		}
	}
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
