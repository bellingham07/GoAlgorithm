package leetcode

import (
	"math"
)

// https://leetcode.cn/problems/minimum-window-substring/

//给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。

//注意：

//对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
//如果 s 中存在这样的子串，我们保证它是唯一的答案。

// 示例 1：
//
// 输入：s = "ADOBECODEBANC", t = "ABC"
// 输出："BANC"
// 解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
func minWindow(s string, t string) string {
	ori, cnt := map[byte]int{}, map[byte]int{}
	// 统计t里面出现的字符以及其次数
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	sLen := len(s)
	// 最长的长度
	leng := math.MaxInt32
	ansL, ansR := -1, -1

	// 检查如果t（ori）里面的字符全部出现在cnt数组里面，说明已经找到终点，此时需要开始移动起点，确定最小长度
	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}

	for l, r := 0, 0; r < sLen; r++ {
		// 当前元素是我们需要的元素 将其放入cnt中
		if ori[s[r]] > 0 {
			cnt[s[r]]++
		}
		// 判断是否找到终点，以及是否越界
		for check() && l <= r {
			// 找到终点
			// 如果当前长度小于已知长度
			if r-l+1 < leng {
				leng = r - l + 1
				// 修改左右边界值
				ansL, ansR = l, l+leng
			}
			// 字符是目标字符
			// 将该目标字符从临时的移除，使其跳出此循坏
			if _, ok := ori[s[l]]; ok {
				cnt[s[l]] -= 1
			}
			// 移动开始位置
			l++
		}
	}
	// 数值没有改变说明没有符合要求 返回""
	if ansL == -1 {
		return ""
	}
	// 截取目标字符串并返回
	return s[ansL:ansR]
}

func minWindow2(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	// 目标数组
	m := map[byte]int{}
	// 临时数组
	temp := map[byte]int{}
	for i := range t {
		m[t[i]]++
	}

	left := -1
	right := 0
	length := len(s) + 1

	for i, j := 0, 0; j < len(s); j++ {
		if m[s[j]] > 0 {
			temp[s[j]]++
		}
		for check(s[i:j], m) && i < j {
			subL := j - i + 1
			if subL < length {
				length = subL
				left = i
				right = left + length
			}
			if _, ok := m[s[i]]; ok {
				temp[s[i]] -= 1
			}
			i++
		}
	}
	if left == -1 {
		return ""
	}
	return s[left:right]
}

func check(curStr string, m map[byte]int) bool {
	for i := 0; i < len(curStr); i++ {
		if _, ok := m[curStr[i]]; ok {
			// exist
			m[curStr[i]]--
		} else {
			continue
		}
		if m[curStr[i]] == 0 {
			delete(m, curStr[i])
		}
	}
	if len(m) == 0 {
		return true
	}
	return false
}
