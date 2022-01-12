package slicetest

func SumAll(numbersAll ...[]int) []int {
	//lenOfNumbers := len(numbersAll) // 参数的长度
	//sums := make([]int, lenOfNumbers)
	//for i, numbers := range numbersAll {
	//	sums[i] = Sum(numbers)
	//}
	// 重构
	var sums []int
	for _, numbers := range numbersAll {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func Sum(numbers []int) int {
	//
	var sum int
	for _, v := range numbers {
		sum += v
	}
	return sum
}
