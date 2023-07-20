package leetcode

func generateMatrix(n int) [][]int {
	loop := n / 2
	center := n / 2
	offset := 1
	x := 0
	y := 0
	i, j := 0, 0

	ans := make([][]int, n)
	for i = 0; i < n; i++ {
		ans[i] = make([]int, n)
	}

	count := 1
	for loop > 0 {
		i, j = x, y
		for j = y; j < n-offset; j++ {
			ans[i][j] = count
			count++
		}
		for i = x; i < n-offset; i++ {
			ans[i][j] = count
			count++
		}
		for ; j > y; j-- {
			ans[i][j] = count
			count++
		}
		for ; i > x; i-- {
			ans[i][j] = count
			count++
		}

		x++
		y++
		offset++
		loop--
	}
	if n%2 == 1 {
		ans[center][center] = n * n
	}
	return ans
}
