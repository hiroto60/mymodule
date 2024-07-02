package mathutils

import "testing"

func TestAdd(t *testing.T) {

	testCases := []struct{ a, b, expected int }{
		{1, 2, 3},
		{0, 1, 1},
		{-1, 1, 0},
		{-1, -1, -2},
	}

	for _, tc := range testCases {
		if Add(tc.a, tc.b) != tc.expected {
			t.Errorf("Expected %d + %d to equal %d", tc.a, tc.b, tc.expected)
		}
	}

}

func TestSubtract(t *testing.T) {

	testCases := []struct{ a, b, expected int }{
		{2, 1, 1},
		{0, 1, -1},
		{-3, 1, -4},
		{-1, -1, 0},
	}

	for _, tc := range testCases {
		if Subtract(tc.a, tc.b) != tc.expected {
			t.Errorf("Expected %d - %d to equal %d", tc.a, tc.b, tc.expected)
		}
	}
}
