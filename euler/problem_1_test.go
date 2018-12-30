package eulerproblems

import "testing"

/**
 * Given a number, sum all the multiples of 3 or 5 below it.
 *  Example: given 10, the multiples are 3, 5, 6, and 9; the sum of which are 23.
 *
 *  Tasks:
 *   - Given a number, construct a slice containing all multiples of 3 or 5 up to given number
 *   - Given a slice of numbers, return their sum
 */

func TestProblemOne(t *testing.T) {
	actual := problemOne(1000)
	expected := 233168

	if actual != expected {
		t.Fatalf("Got: %d, expected: %d", actual, expected)
	}
}

func TestBuildMultiplesSlice(t *testing.T) {
	actual := buildMultiplesSlice(10)
	expected := []int{3, 5, 6, 9}

	if !areEqual(actual, expected) {
		t.Fatalf("Got: %v, expected: %v", actual, expected)
	}
}

func TestSumSlice(t *testing.T) {
	actual := sumSlice([]int{3, 5, 6, 9})
	expected := 23

	if actual != expected {
		t.Fatalf("Got: %d, expected: %d", actual, expected)
	}
}

// TODO: Make utility package?
func areEqual(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, num := range a {
		if num != b[i] {
			return false
		}
	}

	return true
}
