package leetcode

func removeComments(source []string) []string {
	res := []string{}
	new_line := []byte{}
	in_block := false
	for _, line := range source {
		for i := 0; i < len(line); i++ {
			if in_block {
				if i+1 < len(line) && line[i] == '*' && line[i+1] == '/' {
					in_block = false
					i++
				}
			} else {
				if i+1 < len(line) && line[i] == '/' && line[i+1] == '*' {
					in_block = true
					i++
				} else if i+1 < len(line) && line[i] == '/' && line[i+1] == '/' {
					break
				} else {
					new_line = append(new_line, line[i])
				}
			}
		}
		if !in_block && len(new_line) > 0 {
			res = append(res, string(new_line))
			new_line = []byte{}
		}
	}
	return res
}
