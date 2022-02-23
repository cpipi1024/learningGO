package main

import "fmt"

func solution_dongtaiguihua(nums []int) int {
	// 动态规划
	// 状态:dp[i] = max(dp[i-1], dp[i-1] + nums[i])			阶段：下标走到i的数组
	// 决策：if dp[i-1] < 0{ dp[i] = nums[i]}else{dp[i-1] + nums[i]}
	// 边界条件：i > 1

	if len(nums) == 1 {
		return nums[0]
	}
	//pre_max := nums[0]
	//now_max := 0
	var max int = nums[0]
	dp_array := make([]int, len(nums), 2*len(nums))
	dp_array[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp_array[i-1]+nums[i] < nums[i] {
			dp_array[i] = nums[i]
		} else {
			dp_array[i] = nums[i] + dp_array[i-1]
		}

	}
	for _, v := range dp_array {
		// 负数之间的大小比较
		if v > max {
			max = v
		}
	}
	return max
}

func solution_tanxin(nums []int) int {
	// 贪心法
	if len(nums) == 1 {
		return nums[0]
	}

	now_max := 0
	ture_max := nums[0]
	pre_max := nums[0]
	for i := 1; i < len(nums); i++ {
		if pre_max < 0 {
			now_max = nums[i]
		} else {
			now_max = pre_max + nums[i]
		}
		if now_max > ture_max {
			ture_max = now_max
		}
		pre_max = now_max
	}
	return ture_max
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	res := solution_tanxin(nums)
	fmt.Printf("最大子数组和%d \n", res)
}
