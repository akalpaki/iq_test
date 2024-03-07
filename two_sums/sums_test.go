package sums

import (
	"testing"
)

func TestTwoSumsMyAlgo(t *testing.T) {
	type testCase struct {
		name     string
		data     []int
		target   int
		expected []int
	}

	tc := []testCase{
		{
			name:     "example 1",
			data:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "example 2",
			data:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			name:     "example 3",
			data:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
	}

	for _, tt := range tc {
		res := twoSumMyAlgo(tt.data, tt.target)
		if res[0] != tt.expected[0] && res[1] != tt.expected[1] {
			t.Errorf("Test failed on case %s: expected=%d, result=%d", tt.name, tt.expected, res)
		}
	}
}

func BenchmarkTwoSumsMyAlgo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		twoSumMyAlgo([]int{2, 15, 11, 7}, 9)
	}
}

func TestTwoSumsIteration1(t *testing.T) {
	type testCase struct {
		name     string
		data     []int
		target   int
		expected []int
	}

	tc := []testCase{
		{
			name:     "example 1",
			data:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "example 2",
			data:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			name:     "example 3",
			data:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
	}

	for _, tt := range tc {
		res := twoSumIteration1(tt.data, tt.target)
		if res[0] != tt.expected[0] && res[1] != tt.expected[1] {
			t.Errorf("Test failed on case %s: expected=%d, result=%d", tt.name, tt.expected, res)
		}
	}
}

func BenchmarkTwoSumsIteration1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		twoSumIteration1([]int{2, 15, 11, 7}, 9)
	}
}

func TestTwoSumsIteration2(t *testing.T) {
	type testCase struct {
		name     string
		data     []int
		target   int
		expected []int
	}

	tc := []testCase{
		{
			name:     "example 1",
			data:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "example 2",
			data:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			name:     "example 3",
			data:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
	}

	for _, tt := range tc {
		res := twoSumIteration2(tt.data, tt.target)
		if res[0] != tt.expected[0] && res[1] != tt.expected[1] {
			t.Errorf("Test failed on case %s: expected=%d, result=%d", tt.name, tt.expected, res)
		}
	}
}

func BenchmarkTwoSumsIteration2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		twoSumIteration2([]int{2, 15, 11, 7}, 9)
	}
}
