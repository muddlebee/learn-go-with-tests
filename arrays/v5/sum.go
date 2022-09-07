package main

// Sum calculates the total from a slice of numbers.
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll calculates the respective sums of every slice passed in.
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	//use append keyword for dynamic arrays
	// for _, numbers := range numbersToSum {
	// 	sums = append(sums, Sum(numbers))
	// }

	for i := 0; i < len(numbersToSum); i++ {
		sums = append(sums, Sum(numbersToSum[i]))

	}

	return sums
}
