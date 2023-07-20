package offer

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {

	}
	n := len(matrix)
	n1 := len(matrix[0])
	loop := n1 / 2
	center := n / 2
	offset := 1
	x := 0
	y := 0
	i, j := 0, 0

	ans := make([]int, n*n1)

	count := 0
	for loop > 0 {
		i, j = x, y
		for j = y; j < n1-offset; j++ {
			ans[count] = matrix[i][j]
			count++
		}
		for i = x; i < n-offset; i++ {
			ans[count] = matrix[i][j]
			count++
		}
		for ; j > y; j-- {
			ans[count] = matrix[i][j]
			count++
		}
		for ; i > x; i-- {
			ans[count] = matrix[i][j]
			count++
		}

		x++
		y++
		offset++
		loop--
	}
	if n%2 == 1 && n1%2 == 1 {
		ans[count] = matrix[center][center]
	}
	return ans
}
