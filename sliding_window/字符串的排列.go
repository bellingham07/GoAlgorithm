package sliding_window

//https://leetcode.cn/problems/permutation-in-string/

func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}
	// 定义两个map 一个负责存储s1中字符出现的次数 一个负责临时记录字符出现次数
	ori := map[byte]int{}
	cnt := map[byte]int{}
	for i := 0; i < len(s1); i++ {
		ori[s1[i]]++
	}

	// 判断s2中是否已经全部出现s1的字符
	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}

	j := 0 // 起始位置
	i := 0 // 结束位置
	l := len(s1)
	// 滑动窗口
	for ; i < len(s2); i++ {
		// 将s2中出现的s1中的字符存入cnt
		if ori[s2[i]] > 0 {
			cnt[s2[i]]++
		}
		// 是否满足条件 满足条件开始 滑动窗口
		for check() && j <= i {
			if cnt[s2[j]] > 0 {
				// 如果出现了第一个字符是我们所需要的字符， 按照题目描述 那么接下来的所有字符都应该是符合我们要求的字符，所以只需要直接判断长度是否相等
				// 相等 返回true
				if l == i-j+1 {
					return true
				} else {
					// 不相等 将当前出现的符合要求的字符移除
					// 比如这样的 s := "adc"
					//	s2 := "dcda"
					cnt[s2[j]]--
				}
			}
			// 移动滑动窗口起始位置
			j++
		}
	}
	return false
}

// 以下为leetcode官方

// 双指针
func checkInclusion2(s1, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}
	cnt := [26]int{}
	// 将所需要的字符数量置为负数
	for _, ch := range s1 {
		cnt[ch-'a']--
	}
	left := 0
	for right, ch := range s2 {
		x := ch - 'a'
		cnt[x]++
		// 如果该元素数量大于0 说明不是我们所需要的元素
		for cnt[x] > 0 {
			// 将刚刚放进去的元素减一，保证了目标字符的数量始终为负数
			cnt[s2[left]-'a']--
			// 移动起始位置
			left++
		}
		// 为什么这里长度相等就可以直接返回？
		// 以 s1="ab",s2="eidbaooo" 为例
		// l的位置只会停留在两个地方，一个是起始位置0，一个是符合要求的字符下面，而r会一直移动，当r移动过程中，l起始位置没发生改变。
		// 说明这其中的字符全部符合要求，所以当r-l+1==len（s1）时，我们就找到了目标字符串，可以直接返回true
		// 可以画图试试
		if right-left+1 == n {
			return true
		}
	}
	return false
}

// 滑动窗口
func checkInclusion3(s1, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}
	cnt := [26]int{}
	for i, ch := range s1 {
		cnt[ch-'a']--
		cnt[s2[i]-'a']++
	}
	diff := 0
	for _, c := range cnt[:] {
		if c != 0 {
			diff++
		}
	}
	if diff == 0 {
		return true
	}
	for i := n; i < m; i++ {
		x, y := s2[i]-'a', s2[i-n]-'a'
		if x == y {
			continue
		}
		if cnt[x] == 0 {
			diff++
		}
		cnt[x]++
		if cnt[x] == 0 {
			diff--
		}
		if cnt[y] == 0 {
			diff++
		}
		cnt[y]--
		if cnt[y] == 0 {
			diff--
		}
		if diff == 0 {
			return true
		}
	}
	return false
}
