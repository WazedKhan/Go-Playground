package main

func Sum(numbers []int) int {
	sum := 0

	for _, value := range numbers{
		sum += value
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	res := make([]int, len(numbersToSum))
	for i, numbers := range numbersToSum{
		res[i] = Sum(numbers)
	}
	return res
}