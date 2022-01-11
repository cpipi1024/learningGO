package iteration

//const RepeatTimes = 5

func Repeat(words string, times int) string {
	res := ""
	for i := 0; i < times; i++ {
		res += words
	}
	return res
}
