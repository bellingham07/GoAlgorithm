package offer

import "math"

// https://leetcode.cn/problems/gu-piao-de-zui-da-li-run-lcof/

//输入: [7,1,5,3,6,4]
//输出: 5
//解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
//注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。

func maxProfit(prices []int) int {
	ans := 0
	length := len(prices)

	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if prices[j]-prices[i] > ans {
				ans = prices[j] - prices[i]
			}
		}
	}
	return ans
}

func maxProfit2(prices []int) int {
	max := 0
	min := math.MaxInt

	// 只需要记录我们卖出之前的最低价格就行了
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else if prices[i]-min > max {
			max = prices[i] - min
		}
	}
	return max
}
