package main

import "fmt"

func solution_1(prices []int) int {
	var max = func(isTrue bool, a, b int) int {
		if isTrue {
			return a
		}
		return b
	}
	// prices[i]代表股票的在i天的价格
	L := len(prices) //股票根据天数变化的价格
	max_profit := 0  // 最大收益
	//profit_array := make([]int, L)
	now_profit := 0 // dp[i]
	pre_profit := 0 // dp[i-1]
	min := prices[0]
	// 转移方程dp[i] = Max{dp[i]-min(dp[i-1]), dp[i-1]}
	for i := 1; i < L; i++ {
		//先找出最小值
		if prices[i] < min {
			min = prices[i]
		}
		now_profit = prices[i] - min
		max_profit = max(now_profit > pre_profit, now_profit, pre_profit)
		pre_profit = max_profit
	}
	if max_profit < 0 {
		return 0
	}
	return max_profit
}

/* func find_max(profit_array []int) int {
	max := profit_array[0]
	for _, v := range profit_array {
		if v > max {
			max = v
		}
	}
	return max
} */

/* func max(a, b int) int {
	if a > b {
		return a
	}
	return b
} */

func main() {
	prices := []int{7, 6, 4, 3, 1}
	profit := solution_1(prices)
	fmt.Printf("最大利润是:%d\n", profit)
}
