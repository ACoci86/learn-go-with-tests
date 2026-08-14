package grades

import "slices"

func Average(numbers []int) int {

	if len(numbers) == 0 {
		return 0
	}
	sum := 0
	for i := 0; i < len(numbers); i++ {

		sum += numbers[i]
	}
	return sum / len(numbers)
}

func Highest(numbers []int) int {

	if len(numbers) == 0 {
		return 0
	}

	return slices.Max(numbers)

}

func Lowest(numbers []int) int {

	if len(numbers) == 0 {
		return 0
	}

	return slices.Min(numbers)

}

func Passing(numbers []int) []int {
	var passed []int

	for _, number := range numbers {
		if number >= 60 {
			passed = append(passed, number)
		}

	}
	return passed

}
