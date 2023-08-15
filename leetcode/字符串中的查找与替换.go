package leetcode

import "fmt"

func findReplaceString(s string, indices []int, sources []string, targets []string) string {
	ch := []byte(s)
	res := ""
	for i := 0; i < len(ch); {
		idx := in(i, indices)
		if idx < len(indices) && idx != -1 {
			if string(ch[indices[idx]:indices[idx]+len(sources[idx])]) == sources[idx] {
				res += targets[idx]
				i += len(sources[idx])
			} else {
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

func in(idx int, arr []int) int {
	for i := 0; i < len(arr); i++ {
		if idx == arr[i] {
			return i
		}
	}
	return -1
}
