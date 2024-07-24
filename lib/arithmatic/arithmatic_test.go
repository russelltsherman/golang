package arithmatic

import "testing"

func TestSubtract(t *testing.T) {
	a := 5
	b := 6
	expected := -1
	result := Subtract(a, b)
	if result != expected {
		t.Errorf("Subtract function failed expected %d received %d", expected, result)
	}

}

func TestSum(t *testing.T) {
	a := 5
	b := 6
	expected := 11
	result := Sum(a, b)
	if result != expected {
		t.Errorf("Sum function failed expected %d received %d", expected, result)
	}
}

func TestMultiply(t *testing.T) {
	a := 5
	b := 6
	expected := 30
	result := Multiply(a, b)
	if result != expected {
		t.Errorf("multiply function failed expected %d received %d", expected, result)
	}

}
