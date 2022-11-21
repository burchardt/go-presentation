package my_math

import (
	"testing"
)

func TestPlus(t *testing.T) {
	actual := Plus(2, 3)
	expected := 5

	if actual != expected {
		t.Errorf("Expected result to be %d but instead got %d!", expected, actual)
	}
}

func TestSum(t *testing.T) {
	actual := Sum(1, 2, 3, 4, 5)
	expected := 15

	if actual != expected {
		t.Errorf("Expected result to be %d but instead got %d!", expected, actual)
	}
}

func TestSumAndAverage(t *testing.T) {
	total, average := SumAndAverage(10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20)
	expectedSum, expectedAverage := 165, 15.0

	if total != expectedSum {
		t.Errorf("Expected result to be %d but instead got %d!", expectedSum, total)
	}

	if average != expectedAverage {
		t.Errorf("Expected result to be %.2f but instead got %.2f!", expectedAverage, average)
	}
}

func TestDivide(t *testing.T) {
	if actual, err := Divide(8, 4); err == nil {
		expected := 2
		if actual != expected {
			t.Errorf("Expected result to be %d but instead got %d!", expected, actual)
		}
	} else {
		t.Errorf("Divide by a non-zero denominator should be successful, but got an error: %s", err)
	}

	if _, err := Divide(1, 0); err == nil {
		t.Errorf("Divide by zero should cause an error, but got nil instead!")
	}
}
