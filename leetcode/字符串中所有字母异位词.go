package leetcode

// https://leetcode.cn/problems/find-all-anagrams-in-a-string/

// 滑动窗口思想

func findAnagrams(s string, p string) []int {
	// 返回结果
	ans := make([]int, 0)
	// 将目标字符全部存入ori
	ori := map[byte]int{}
	// 临时map 用于存储临时字符集
	cnt := map[byte]int{}

	length := len(p)

	// 放入字符
	for i := 0; i < length; i++ {
		ori[p[i]]++
	}

	// 检查两个map是否相等
	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}

	j := 0
	i := 0

	// 滑动窗口
	for ; j < len(s); j++ {
		// 将字符放入临时字符
		cnt[s[j]]++
		// 满足条件 开始滑动窗口
		for check() && i <= j {
			if j-i+1 == length {
				ans = append(ans, i)
			}
			// 将是目标字符但却不符合题目要求的字符移除
			if _, ok := ori[s[i]]; ok {
				cnt[s[i]]--
			}
			i++
		}
	}
	return ans
}
