package dp

func uniquePaths(m int, n int) int {
	// 1.确定dp数组下表及含义
	// 2.确定递推公式
	// 3.初始化dp数组
	// 4.确定遍历顺序
	// 5.举例验证
	dp := make([][]int, m) // 创建一个切片切片，每个切片表示一行

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n) // 为每行创建一个切片，表示列
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := make([][]int, m) // 创建一个切片切片，每个切片表示一行

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n) // 为每行创建一个切片，表示列
	}

	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}
		dp[i][0] = 1
	}

	for i := 0; i < n; i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}
		dp[0][i] = 1

	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}

		}
	}
	return dp[m-1][n-1]
}
