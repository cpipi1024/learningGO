package main

func solution(nums []int) bool {
	// 创建映射对象
	var temp_map = make(map[int]int)
	for _, value := range nums {
		_, ok := temp_map[value]
		if ok {
			return true
		}
		temp_map[value] = 1
	}
	return false
}

/* func main() {
	nums := []int{1, 2, 3, 4, 5}
	res := solution(nums)
	fmt.Printf("是否存在重复数字 %v\n", res)
}
*/
