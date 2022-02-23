package main

type sum int

func solution_2(nums []int) sum {
	// 数组长度为1的情况
	if len(nums) == 1 {
		return sum(nums[0])
	}
	temp_sum := nums[0]
	// 动态规划  最小规划数从1开始
	for i := 1; i <= len(nums); i++ {
		// row
		// 越界问题
		for j := 0; j <= len(nums)-i; j++ {
			// col
			ttemp_sum := sumall(nums[j : j+i])
			if ttemp_sum > temp_sum {
				temp_sum = ttemp_sum
			}
		}
	}
	return sum(temp_sum)
}

func sumall(nums []int) int {
	tot := 0
	for _, val := range nums {
		tot += val
	}
	return tot
}

/* func main() {
	//nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	nums := []int{5, 4, -1, 7, 8}
	//nums := []int{-2, -1}
	sum := solution_2(nums)
	fmt.Printf("该数组的最大子数组和是:%d\n", sum)
}
*/
