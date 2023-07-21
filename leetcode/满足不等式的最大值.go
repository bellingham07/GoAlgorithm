package leetcode

import (
	"math"
)

//给你一个数组 points 和一个整数 k 。数组中每个元素都表示二维平面上的点的坐标，并按照横坐标 x 的值从小到大排序。也就是说 points[i] = [xi, yi] ，
//并且在 1 <= i < j <= points.length 的前提下， xi < xj 总成立。
//请你找出 yi+ yj+ |xi- xj| 的 最大值，其中 |xi- xj|<= k 且 1 <= i < j <= points.length。
//题目测试数据保证至少存在一对能够满足 |xi- xj|<= k 的点。

//输入：points = [[1,3],[2,0],[5,10],[6,-10]], k = 1
//输出：4
//解释：前两个点满足 |xi- xj| <= 1 ，代入方程计算，则得到值 3 + 0 + |1 - 2| = 4 。第三个和第四个点也满足条件，得到值 10 + -10 + |5 - 6| = 1 。
//没有其他满足条件的点，所以返回 4 和 1 中最大的那个。
//链接：https://leetcode.cn/problems/max-value-of-equation

/*
f(x)=yi+ yj+ |xi- xj| 因为x是递增数列 所以xi-xj恒小于0 所以去掉绝对符号
	=yi+ yj-（xi- xj）
	=（yi-xi）+(yj+xj)

*/

// 滑动窗口会超时，也不算滑动窗口 双循坏会超时
func findMaxValueOfEquation(points [][]int, k int) int {
	abs := func(i, j int) int {
		t := i - j
		if t < 0 {
			return -t
		}
		return t
	}

	calculate := func(x1, y1, x2, y2 int) int {
		//yi+ yj+ |xi- xj|
		return y1 + y2 + abs(x1, x2)
	}

	if len(points) < 3 {
		return calculate(points[0][0], points[0][1], points[1][0], points[1][1])
	}

	maxValue := -math.MaxInt
	for j := 1; j < len(points); j++ {
		for p := 0; p < j; p++ {
			if abs(points[p][0], points[j][0]) <= k {
				temp := calculate(points[p][0], points[p][1], points[j][0], points[j][1])
				if maxValue < temp {
					maxValue = temp
				}
			}
		}
	}

	return maxValue
}

func findMaxValueOfEquation3(points [][]int, k int) int {
	res := -0x3f3f3f3f
	q := [][]int{}
	for _, point := range points {
		x, y := point[0], point[1]
		for len(q) > 0 && x-q[0][1] > k {
			q = q[1:]
		}
		if len(q) > 0 {
			res = max(res, x+y+q[0][0])
		}
		for len(q) > 0 && y-x >= q[len(q)-1][0] {
			q = q[:len(q)-1]
		}
		q = append(q, []int{y - x, x})
	}
	return res
}

func findMaxValueOfEquation2(points [][]int, k int) int {
	type s struct{ a, b int }
	ans := math.MinInt32
	q := []s{{points[0][0], points[0][1] - points[0][0]}}
	for i := 1; i < len(points); i++ {
		for len(q) > 0 && (points[i][0]-q[0].a) > k {
			q = q[1:]
		}
		if len(q) > 0 {
			ans = max2(ans, q[0].b+points[i][0]+points[i][1])
		}
		for len(q) > 0 && (q[len(q)-1].b) < (points[i][1]-points[i][0]) {
			q = q[:len(q)-1]
		}
		q = append(q, s{points[i][0], points[i][1] - points[i][0]})
	}
	return ans

}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}
