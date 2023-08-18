package leetcode

import "fmt"

func findReplaceString(s string, indices []int, sources []string, targets []string) string {
	ch := []byte(s)
	res := ""
	// 遍历每一个字符串看是否需要被换
	for i := 0; i < len(ch); {
		idx := in(i, indices)
		// 当前位置是否有修改数据
		if idx < len(indices) && idx != -1 {
			// 判断目标字符串是否存在于源字符串中
			if string(ch[indices[idx]:indices[idx]+len(sources[idx])]) == sources[idx] {
				res += targets[idx]
				i += len(sources[idx])
			} else {
				// 避免字符串添加过长，出现越界现象
				if len(sources[idx]) > len(ch)-indices[idx] {
					res += string(ch[indices[idx]:])
				} else {
					res += string(ch[indices[idx] : indices[idx]+len(sources[idx])])
				}
				i += len(sources[idx])
			}
			idx++
		} else {
			res += string(ch[i])
			i++
		}
	}
	c := []byte(res)

	for i, b := range c {
		fmt.Println(i, " ", b)
	}

	return res
}

// 检查当前字符是否在需要被替换的字符中
func in(idx int, arr []int) int {
	for i := 0; i < len(arr); i++ {
		if idx == arr[i] {
			return i
		}
	}
	return -1
}
