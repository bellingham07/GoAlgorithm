package leetcode

// https://leetcode.cn/problems/fruit-into-baskets/

//你正在探访一家农场，农场从左到右种植了一排果树。这些树用一个整数数组 fruits 表示，其中 fruits[i] 是第 i 棵树上的水果 种类 。
//你想要尽可能多地收集水果。然而，农场的主人设定了一些严格的规矩，你必须按照要求采摘水果：
//你只有 两个 篮子，并且每个篮子只能装 单一类型 的水果。每个篮子能够装的水果总量没有限制。
//你可以选择任意一棵树开始采摘，你必须从 每棵 树（包括开始采摘的树）上 恰好摘一个水果 。采摘的水果应当符合篮子中的水果类型。每采摘一次，你将会向右移动到下一棵树，并继续采摘。
//一旦你走到某棵树前，但水果不符合篮子的水果类型，那么就必须停止采摘。
//给你一个整数数组 fruits ，返回你可以收集的水果的 最大 数目。

func totalFruit(fruits []int) int {
	basket := make(map[int]int) // 菜篮子容量 通过map的方式 来记录篮子是否装满 每种装了多少
	left := 0                   // 滑动窗口起始点
	res := 0                    // 这里求的是最长 所以将结果初始化为0

	for right, fruit := range fruits {
		// 将已采摘的放入篮子中
		basket[fruit]++
		// 当篮子中种类多于三个时 开始滑动
		for len(basket) > 2 {
			// 将fruits[left]一个一个从map中去除
			f := fruits[left]
			basket[f]--
			// 当这种水果数量为0时候，将其从map中去除
			if basket[f] == 0 {
				delete(basket, f)
			}
			left++
		}
		// 记录下当前终点时 所有的最长长度
		subL := right - left + 1
		if res < subL {
			res = subL
		}
	}
	return res
}
