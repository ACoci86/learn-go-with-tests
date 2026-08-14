package main

import (
	"fmt"

	"learning-golang/projects/grades"
)

func main() {
	scores := []int{60, 75, 25, 45, 96, 74, 3}

	fmt.Println("Scores: ", scores)
	fmt.Println("Average:", grades.Average(scores))
	fmt.Println("Highest:", grades.Highest(scores))
	fmt.Println("Lowest: ", grades.Lowest(scores))
	fmt.Println("Passed: ", grades.Passing(scores))
}
