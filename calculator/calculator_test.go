package calculator_test

import (
	"calculator"
	"math"
	"testing"
)

type testCase struct {
	a, b float64
	want float64
}

func closeEnough(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestAdd(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 2, want: 3},
		{a: 0, b: 0, want: 0},
		{a: -1, b: -1, want: -2},
		{a: -1, b: 1, want: 0},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}

	var want float64 = 4
	got := calculator.Add(2, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 2, b: 2, want: 0},
		{a: 1, b: 2, want: -1},
		{a: 0, b: 0, want: 0},
		{a: -1, b: -1, want: 0},
		{a: -1, b: 1, want: -2},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Subtract(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}

	var want float64 = 2
	got := calculator.Subtract(4, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 2, want: 2},
		{a: 0, b: 0, want: 0},
		{a: -1, b: -1, want: 1},
		{a: -1, b: 1, want: -1},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Multiply(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}

	var want float64 = 4
	got := calculator.Multiply(2, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 2, b: 2, want: 1},
		{a: -1, b: -1, want: 1},
		{a: 10, b: 2, want: 5},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		if err != nil {
			t.Errorf("Divide(%f, %f): unexpected error %v", tc.a, tc.b, err)
		}

		if !closeEnough(tc.want, got, 0.001) {
			t.Errorf("Add(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}

	var want float64 = 4
	got := calculator.Add(2, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestDivideInvalid(t *testing.T) {
	t.Parallel()

	_, err := calculator.Divide(2, 0)
	if err == nil {
		t.Errorf("Divide(%f, %f): expected an error but did not get one", 2.0, 0.0)
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 4, b: 0, want: 2},
		{a: 1, b: 0, want: 1},
		{a: 0, b: 0, want: 0},
		{a: 9, b: 0, want: 3},
	}

	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.a)
		if err != nil {
			t.Errorf("Sqrt(%f): unexpected error %v", tc.a, err)
		}

		if !closeEnough(tc.want, got, 0.001) {
			t.Errorf("Sqrt(%f): want %f, got %f", tc.a, tc.want, got)
		}
	}

}

func TestSqrtInvalid(t *testing.T) {
	t.Parallel()

	_, err := calculator.Sqrt(-1)
	if err == nil {
		t.Errorf("Sqrt(%f): expected an error but did not get one", -1.0)
	}
}
