package leetcode

//https://leetcode.cn/problems/lemonade-change/

func lemonadeChange(bills []int) bool {
	m1 := map[int]int{}
	for i := 0; i < len(bills); i++ {
		if bills[i] == 5 {
			// 5
			m1[bills[i]]++
			continue
		} else if bills[i] == 10 {
			// 10
			if m1[5] > 0 {
				m1[5]--
				m1[10]++
				continue
			} else {
				return false
			}
		} else {
			// 20
			if m1[10] > 0 && m1[5] > 0 {
				m1[10]--
				m1[5]--
				m1[20]++
				continue
			}
			if m1[5] > 2 {
				m1[5] -= 3
				m1[20]++
				continue
			}
			return false
		}
	}
	return true
}

func lemonadeChange2(bills []int) bool {
	// 我们发现不管什么情况，我们的手中都必须有5美元。否则下一次的购买就无法找零
	five, ten := 0, 0
	for i := 0; i < len(bills); i++ {
		if bills[i] == 5 {
			five++
		} else if bills[i] == 10 {
			five--
			ten++
		} else if ten > 0 {
			ten--
			five--
		} else {
			five -= 3
		}
		if five < 0 {
			return false
		}
	}
	return true
}
