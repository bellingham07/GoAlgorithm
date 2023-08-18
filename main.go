package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 8; i++ {
		fmt.Println(i % 7)
	}
	fmt.Println(circularGameLosers(5, 2))
}

func circularGameLosers(n int, k int) []int {
	var ans []int
	gamers := make([]int, n) //记录所有玩家的得分，其中第i个玩家的编号为i-1（从0开始）
	cur := 0                 //当前位置
	gamers[cur]++
	for i := 1; ; i++ { //当前次数为第i次
		cur = (cur + i*k) % n
		gamers[cur]++
		if gamers[cur] >= 2 { //当某个朋友第 2 次接到球时，游戏结束
			break
		}
	}
	for i, v := range gamers {
		if v == 0 {
			ans = append(ans, i+1)
		}
	}
	return ans
}
