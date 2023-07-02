package multiply

import (
	"testing"
)

type TestCase struct {
	a, b     int
	expected int
}

// table-driven testcase
func TestMultiply(t *testing.T) {
	testCases := []TestCase{
		{2, 3, 6},
		{4, 5, 20},
		{0, 10, 0},
		{-1, 20, -20},
	}

	for _, tc := range testCases {
		result := Multiply(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("error occurs! a: %d, b: %d, expected: %d", tc.a, tc.b, tc.expected)
		}
	}
}
