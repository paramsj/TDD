package arrays

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum = sum + number
	}
	return sum
}

func DirtySumAll(number ...[]int) int {
	sum := 0

	for _, row := range number {
		for _, element := range row {
			sum = sum + element
		}
	}

	return sum
} // this works but maybe clean it up?

func SumAll(number ...[]int) []int {
	length := len(number)
	sums := make([]int, length)

	for i, number := range number {
		sums[i] = Sum(number)
	}

	return sums
}
