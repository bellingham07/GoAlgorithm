package backTrack

import "sort"

var (
	res  [][]int
	path []int
)

func combinationSum(candidates []int, target int) [][]int {
	res, path := make([][]int, 0), make([]int, 0, len(candidates))
	// 排序是为了方便剪枝
	sort.Ints(candidates)
	dfs(candidates, 0, target)
	return res
}

func dfs(candidates []int, start int, target int) {
	// target==0说明找到符合条件的值
	if target == 0 {
		// 因为没有对path数组进行一个新的扩容，所以地址不会有变化，所以在这里要使用copy函数，不然，直接append，res的指针会一直指向path
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		return
	}
	for i := start; i < len(candidates); i++ {
		// 当前分支已经不满足条件，剪枝
		if candidates[i] > target {
			break
		}
		path = append(path, candidates[i])
		// 选取可重复的，所以此处直接传i
		// 选取不可重复的，此处要传i+1
		dfs(candidates, i, target-candidates[i])
		path = path[:len(path)-1]
	}
}
