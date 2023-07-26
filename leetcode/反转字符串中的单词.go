package leetcode

func reverseWords(s string) string {
	c := []byte(s)
	res := make([]byte, 0)
	for j := len(c) - 1; j > -1; j-- {
		if c[j] != ' ' {
			for i := j; i > -1; i-- {
				if c[i] == ' ' && i != 0 {
					res = append(res, c[i:j+1]...)
					j = i
					break
				}
				if i == 0 {
					res = append(res, ' ')
					res = append(res, c[i:j+1]...)
					j = -1
				}
			}
		}
	}
	return string(res[1:])
}
