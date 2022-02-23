package main

const maxlen = 100

func solution_4(nums []int) int {
	// 求出数组中最长的从左到右递增子序列长度 // 返回最长递增子序的长度
	// nums[1,5,2,4,3]

	max_array := make([]int, 10, maxlen)
	maxv := 0
	for i := len(nums) - 1; i >= 0; i-- {
		for j := i; j < len(nums); j++ {
			if nums[j] > nums[i] {
				max_array[i] = max(max_array[i], max_array[j]+1)
			}
		}
	}
	for _, v := range max_array {
		if v > maxv {
			maxv = v
		}
	}
	return maxv + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/* func main() {
	nums := []int{1, 7, 9, 7, 8, 9, 2, 3, 4}
	maxL := solution_4(nums)
	fmt.Printf("数组中最长的递增子序列长度为%d\n", maxL)
} */
