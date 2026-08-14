package grades

import (
	"fmt"
	"slices"
	"testing"
)

func TestAverage(t *testing.T) {

	scores := []int{10, 20, 30}

	got := Average(scores)
	want := 20

	if got != want {
		t.Errorf("got %d want %d, given %v", got, want, scores)
	}
}

func TestAverageOfEmptyList(t *testing.T) {
	got := Average([]int{})
	want := 0

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestHighest(t *testing.T) {
	scores := []int{45, 64, 1239, 6, 9863}

	got := Highest(scores)
	want := 9863

	if got != want {
		t.Errorf("got %d want %d, given %v", got, want, scores)
	}

	fmt.Println("The highest number is", Highest(scores))

}

func TestLowest(t *testing.T) {
	scores := []int{45, 64, 1239, 6, 9863}

	got := Lowest(scores)
	want := 6

	if got != want {
		t.Errorf("got %d want %d, given %v", got, want, scores)
	}

	fmt.Println("The lowest number is", Lowest(scores))

}

func TestPassing(t *testing.T) {

	scores := []int{60, 75, 25, 45, 96, 74, 3}

	got := Passing(scores)
	want := []int{60, 75, 96, 74}

	if !slices.Equal(got, want) {
		t.Errorf("got %d want %d, given %v", got, want, scores)
	}

	fmt.Println("The passing numbers are", Passing(scores))

}
