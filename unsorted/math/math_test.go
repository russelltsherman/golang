// +build all_tests algorythm_tests

package math

import "testing"

// Here's a unit test to check addition:
// func TestSum(t *testing.T) {
// 	total := Sum(5, 5)
// 	if total != 10 {
// 		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
// 	}
// }

// Test tables
// The concept of "test tables" is a set (slice array) of test inputs and outputs
// Here is an example for the Sum function:

type sumTest struct {
	x        int
	y        int
	expected int
}

var sumTests = []sumTest{
	{1, 1, 2},
	{1, 2, 3},
	{2, 2, 4},
	{5, 2, 7},
	{5, 5, 10},
}

func TestSum(t *testing.T) {
	for _, sumTest := range sumTests {
		total := Sum(sumTest.x, sumTest.y)
		if total != sumTest.expected {
			t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.", sumTest.x, sumTest.y, total, sumTest.expected)
		}
	}
}

type fibTest struct {
	n        int // input
	expected int // expected result
}

var fibTests = []fibTest{
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
	{7, 13},
}

func TestFib(t *testing.T) {
	for _, fibTest := range fibTests {
		actual := Fib(fibTest.n)
		if actual != fibTest.expected {
			t.Errorf("Fib(%d): expected %d, actual %d", fibTest.n, fibTest.expected, actual)
		}
	}
}
