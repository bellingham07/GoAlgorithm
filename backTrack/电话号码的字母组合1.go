package backTrack

var (
	m    []string
	path []byte
	res  []string
)

func letterCombinations(digits string) []string {
	m = []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	path, res = make([]byte, 0), make([]string, 0)
	if digits == "" {
		return res
	}
	dfsletter(digits, 0)
	return res
}

func dfsletter(digits string, start int) {
	// 传进来的字符串长度，就是我们所需要的最终结果的长度
	if len(path) == len(digits) {
		tmp := string(path)
		// 将满足条件的结果放入到结果集中
		res = append(res, tmp)
		return
	}
	// 获取当前我们所要的字母在m切片中的映射位置
	digit := int(digits[start] - '2')
	// 获取当前数字上的全部字母
	str := m[digit]
	for i := 0; i < len(str); i++ {
		// 将第一个放入到切片中
		path = append(path, str[i])
		// 回调函数，遍历下一个数字
		dfsletter(digits, start+1)
		//  清空切片
		path = path[:len(path)-1]
	}
}
