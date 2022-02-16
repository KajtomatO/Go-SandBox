package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	var tail []int

	for _, numbers := range numbersToSum {
		if len(numbers) > 1 {
			tail = numbers[1:]
		} else {
			tail = numbers
		}
		sums = append(sums, Sum(tail))
	}

	return sums
}
