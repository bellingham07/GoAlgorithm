package leetcode

import "sort"

func deleteGreatestValue(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])
	for i := 0; i < row; i++ {
		sort.Ints(grid[i])
	}
	res := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for j := 0; j < col; j++ {
		t := 0
		for i := 0; i < row; i++ {
			t = max(t, grid[i][j])
		}
		res += t
	}
	return res
}
