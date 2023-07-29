package offer

func reverseLeftWords(s string, n int) string {
	c := []byte(s)
	t := make([]byte, 0)
	t = append(t, c[n:]...)
	t = append(t, c[:n]...)
	return string(t)
}
