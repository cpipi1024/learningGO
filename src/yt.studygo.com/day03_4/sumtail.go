package slicetest2

func SumTail(numbersAll ...[]int) []int {
	var sums []int
	for _, numers := range numbersAll {
		sums = append(sums, Sum(numers))
	}
	return sums
}

func Sum(numbers []int) int {
	var sum int
	if len(numbers) == 0 {
		return 0
	}
	for _, number := range numbers[1:] {
		sum += number
	}
	return sum
}
