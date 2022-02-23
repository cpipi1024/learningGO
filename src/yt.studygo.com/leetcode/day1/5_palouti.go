package main

// 一共有M级台阶，每次能走1或2级台阶，请问走上M级台阶需要多少种走法

func solution_5(stairs int) int {
	step_stairs := []int{0, 1, 1}

	for i := 3; i <= stairs; i++ {
		temp := step_stairs[i-1] + step_stairs[i-2] // 状态方程
		step_stairs = append(step_stairs, temp)
	}
	return step_stairs[stairs]
}
func solution_5_1(stairs int) int { // 迭代交换s1 s2 空间优化
	if stairs <= 3 {
		return stairs - 1
	}
	s1 := 1
	s2 := 2
	sum := 0
	for i := 4; i <= stairs; i++ {
		sum = s1 + s2
		s1 = s2
		s2 = sum
	}
	return sum
}

/* func main() {
	stairs := 10
	res := solution_5(stairs)
	res1 := solution_5_1(stairs)
	fmt.Printf("爬上%d级台阶共有%d种方式\n", stairs, res)
	fmt.Printf("爬上%d级台阶共有%d种方式\n", stairs, res1)
} */
